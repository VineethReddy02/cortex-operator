package cortex

import (
	"context"
	cortexv1alpha1 "github.com/VineethReddy02/goModules/cortex-operator/pkg/apis/cortex/v1alpha1"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_cortex")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new Cortex Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileCortex{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("cortex-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource Cortex
	err = c.Watch(&source.Kind{Type: &cortexv1alpha1.Cortex{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner Cortex
	err = c.Watch(&source.Kind{Type: &v1.Deployment{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &cortexv1alpha1.Cortex{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileCortex implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileCortex{}

// ReconcileCortex reconciles a Cortex object
type ReconcileCortex struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a Cortex object and makes changes based on the state read
// and what is in the Cortex.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileCortex) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling Cortex")

	// Fetch the Cortex instance
	instance := &cortexv1alpha1.Cortex{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Check if the deployment already exists, if not create a new one
	found := &v1.Deployment{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {

		err = r.createNginx(instance)
		if err != nil {
			return reconcile.Result{}, err
		}

		err = r.createQueryFrontEnd(instance)
		if err != nil {
			return reconcile.Result{}, err
		}

		err = r.createConsul(instance)
		if err != nil {
			return reconcile.Result{}, err
		}

		err = r.createDistributor(instance)
		if err != nil {
			return reconcile.Result{}, err
		}

		err = r.createIngester(instance)
		if err != nil {
			return reconcile.Result{}, err
		}

		err = r.createS3(instance)
		if err != nil {
			return reconcile.Result{}, err
		}

		err = r.createDynamoDB(instance)
		if err != nil {
			return reconcile.Result{}, err
		}

		err = r.createTableManager(instance)
		if err != nil {
			return reconcile.Result{}, err
		}

		err = r.createConfigsDB(instance)
		if err != nil {
			return reconcile.Result{}, err
		}

		err = r.createConfigs(instance)
		if err != nil {
			return reconcile.Result{}, err
		}

		err =r.createQuerier(instance)
		if err != nil {
			return reconcile.Result{}, err
		}

		// Deployment created successfully - return and requeue
		return reconcile.Result{Requeue: true}, nil
	} else if err != nil {
		reqLogger.Error(err, "Failed to get Deployment")
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil
}


package cortex

import (
	cortexv1alpha1 "github.com/VineethReddy02/goModules/cortex-operator/pkg/apis/cortex/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// deploymentForIngester returns a ingester Deployment object
func (r *ReconcileCortex) serviceForConsul(m *cortexv1alpha1.Cortex) *corev1.Service {
	ls := labelsForConsul()
	service := &corev1.Service{
		TypeMeta:   metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:                      consul,
			Namespace:                 m.Namespace,
		},
		Spec:       corev1.ServiceSpec{
			Ports:  []corev1.ServicePort{{
				Name:       "http",
				Port:       8500,
			}},
			Selector: ls,
		},
		Status:     corev1.ServiceStatus{},
	}

	controllerutil.SetControllerReference(m, service, r.scheme)
	return service
}

// deploymentForIngester returns a ingester Deployment object
func (r *ReconcileCortex) serviceForDistributor(m *cortexv1alpha1.Cortex) *corev1.Service {
	ls := labelsForDistributor()
	service := &corev1.Service{
		TypeMeta:   metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:                      distributor,
			Namespace:                 m.Namespace,
		},
		Spec:       corev1.ServiceSpec{
			Ports:  []corev1.ServicePort{{
				Port:       80,
			}},
			Selector: ls,
		},
		Status:     corev1.ServiceStatus{},
	}

	controllerutil.SetControllerReference(m, service, r.scheme)
	return service
}

// deploymentForIngester returns a ingester Deployment object
func (r *ReconcileCortex) serviceForIngester(m *cortexv1alpha1.Cortex) *corev1.Service {
	ls := labelsForIngester()
	service := &corev1.Service{
		TypeMeta:   metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:                      ingester,
			Namespace:                 m.Namespace,
		},
		Spec:       corev1.ServiceSpec{
			Ports:  []corev1.ServicePort{{
				Port:       80,
			}},
			Selector: ls,
		},
		Status:     corev1.ServiceStatus{},
	}

	controllerutil.SetControllerReference(m, service, r.scheme)
	return service
}

// deploymentForIngester returns a ingester Deployment object
func (r *ReconcileCortex) serviceForQuerier(m *cortexv1alpha1.Cortex) *corev1.Service {
	ls := labelsForQuerier()
	service := &corev1.Service{
		TypeMeta:   metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:                      querier,
			Namespace:                 m.Namespace,
		},
		Spec:       corev1.ServiceSpec{
			Ports:  []corev1.ServicePort{{
				Port:       80,
			}},
			Selector: ls,
		},
		Status:     corev1.ServiceStatus{},
	}

	controllerutil.SetControllerReference(m, service, r.scheme)
	return service
}

func (r *ReconcileCortex) serviceForNginx(m *cortexv1alpha1.Cortex) *corev1.Service {
	ls := labelsForNginx()
	service := &corev1.Service{
		TypeMeta:   metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:                      nginx,
			Namespace:                 m.Namespace,
		},
		Spec:       corev1.ServiceSpec{
			Type:    "NodePort",
			Ports:  []corev1.ServicePort{{
				Name:       "http",
				Port:       80,
				NodePort:   30080,
			}},
			Selector: ls,
		},
		Status:     corev1.ServiceStatus{},
	}

	controllerutil.SetControllerReference(m, service, r.scheme)
	return service
}

func (r *ReconcileCortex) serviceForS3(m *cortexv1alpha1.Cortex) *corev1.Service {
	ls := labelsForS3()
	service := &corev1.Service{
		TypeMeta:   metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:                      s3,
			Namespace:                 m.Namespace,
		},
		Spec:       corev1.ServiceSpec{
			Ports:  []corev1.ServicePort{{
				Port:   4569,
			}},
			Selector: ls,
		},
		Status:     corev1.ServiceStatus{},
	}

	controllerutil.SetControllerReference(m, service, r.scheme)
	return service
}

func (r *ReconcileCortex) serviceForDynamoDb(m *cortexv1alpha1.Cortex) *corev1.Service {
	ls := labelsForDynamo()
	service := &corev1.Service{
		TypeMeta:   metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:                      dynamoDB,
			Namespace:                 m.Namespace,
		},
		Spec:       corev1.ServiceSpec{
			Ports:  []corev1.ServicePort{{
				Port:   8000,
			}},
			Selector: ls,
		},
		Status:     corev1.ServiceStatus{},
	}

	controllerutil.SetControllerReference(m, service, r.scheme)
	return service
}

func (r *ReconcileCortex) serviceForTableManager(m *cortexv1alpha1.Cortex) *corev1.Service {
	ls := labelsForTableManager()
	service := &corev1.Service{
		TypeMeta:   metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:                      tableManager,
			Namespace:                 m.Namespace,
		},
		Spec:       corev1.ServiceSpec{
			Ports:  []corev1.ServicePort{{
				Port:   8000,
			}},
			Selector: ls,
		},
		Status:     corev1.ServiceStatus{},
	}

	controllerutil.SetControllerReference(m, service, r.scheme)
	return service
}

func (r *ReconcileCortex) serviceForConfigs(m *cortexv1alpha1.Cortex) *corev1.Service {
	ls := labelsForConfigs()
	service := &corev1.Service{
		TypeMeta:   metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:                      configs,
			Namespace:                 m.Namespace,
		},
		Spec:       corev1.ServiceSpec{
			Ports:  []corev1.ServicePort{{
				Port:   80,
			}},
			Selector: ls,
		},
		Status:     corev1.ServiceStatus{},
	}

	controllerutil.SetControllerReference(m, service, r.scheme)
	return service
}

func (r *ReconcileCortex) serviceForConfigsDb(m *cortexv1alpha1.Cortex) *corev1.Service {
	ls := labelsForConfigsDB()
	service := &corev1.Service{
		TypeMeta:   metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:                      configsDB,
			Namespace:                 m.Namespace,
		},
		Spec:       corev1.ServiceSpec{
			Ports:  []corev1.ServicePort{{
				Port:   5432,
			}},
			Selector: ls,
		},
		Status:     corev1.ServiceStatus{},
	}

	controllerutil.SetControllerReference(m, service, r.scheme)
	return service
}

func (r *ReconcileCortex) serviceForQueryFrontEnd(m *cortexv1alpha1.Cortex) *corev1.Service {
	ls := labelsForQueryFrontEnd()
	service := &corev1.Service{
		TypeMeta:   metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:                      queryFrontEnd,
			Namespace:                 m.Namespace,
		},
		Spec:       corev1.ServiceSpec{
			ClusterIP:  "None",
			Ports:  []corev1.ServicePort{{
				Name:   "grpc",
				Port:   9095,
			},
			{
				Name:  "http",
				Port:   80,
			}},
			Selector: ls,
		},
		Status:     corev1.ServiceStatus{},
	}

	controllerutil.SetControllerReference(m, service, r.scheme)
	return service
}
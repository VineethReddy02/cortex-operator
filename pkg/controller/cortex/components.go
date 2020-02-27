package cortex

import (
	"context"
	cortexv1alpha1 "github.com/VineethReddy02/goModules/cortex-operator/pkg/apis/cortex/v1alpha1"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/types"
)

const (
	ingester      = "ingester"
	distributor   = "distributor"
	consul        = "consul"
	querier       = "querier"
	queryFrontEnd = "query-frontend"
	s3            = "s3"
	dynamoDB      = "dynamodb"
	nginx         = "nginx"
	alertManager  = "alert-manager"
	ruler         = "ruler"
	configs       = "configs"
	configsDB     = "configs-db"
	tableManager  = "table-manager"
)

func (r *ReconcileCortex) createConsul(instance *cortexv1alpha1.Cortex) error {
	deployment := &v1.Deployment{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: consul, Namespace: instance.Namespace}, deployment)
	if instance.Spec.Consul.Size == 0 || err == nil {
		return nil
	}
	reqLogger := log.WithValues("Request.Namespace", instance.Namespace, "Request.Name", instance.Name)
	consul := r.deploymentForConsul(instance)
	reqLogger.Info("Creating a new Deployment", "Deployment.Namespace", consul.Namespace, "Deployment.Name", consul.Name)
	err = r.client.Create(context.TODO(), consul)
	if err != nil {
		reqLogger.Error(err, "Failed to create new Deployment", "Deployment.Namespace", consul.Namespace, "Deployment.Name", consul.Name)
		return err
	}
	consulService := r.serviceForConsul(instance)
	err = r.client.Create(context.TODO(), consulService)
	if err != nil {
		reqLogger.Error(err, "Failed to create new consul service", "Deployment.Namespace", consul.Namespace, "Deployment.Name", consul.Name)
		return err
	}
	return nil
}

func (r *ReconcileCortex) createDistributor(instance *cortexv1alpha1.Cortex) error {
	deployment := &v1.Deployment{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: distributor, Namespace: instance.Namespace}, deployment)
	if instance.Spec.Distributor.Size == 0 || err == nil {
		return nil
	}
	reqLogger := log.WithValues("Request.Namespace", instance.Namespace, "Request.Name", instance.Name)
	distributor := r.deploymentForDistributor(instance)
	reqLogger.Info("Creating a new Deployment", "Deployment.Namespace", distributor.Namespace, "Deployment.Name", distributor.Name)
	err = r.client.Create(context.TODO(), distributor)
	if err != nil {
		reqLogger.Error(err, "Failed to create new Deployment", "Deployment.Namespace", distributor.Namespace, "Deployment.Name", distributor.Name)
		return err
	}
	distributorService := r.serviceForDistributor(instance)
	err = r.client.Create(context.TODO(), distributorService)
	if err != nil {
		reqLogger.Error(err, "Failed to create new distributor service", "Deployment.Namespace", distributor.Namespace, "Deployment.Name", distributor.Name)
		return err
	}
	return nil
}

func (r *ReconcileCortex) createIngester(instance *cortexv1alpha1.Cortex) error {
	deployment := &v1.Deployment{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: ingester, Namespace: instance.Namespace}, deployment)
	if instance.Spec.Ingester.Size == 0 || err == nil {
		return nil
	}
	reqLogger := log.WithValues("Request.Namespace", instance.Namespace, "Request.Name", instance.Name)
	ingester := r.deploymentForIngester(instance)
	reqLogger.Info("Creating a new Deployment", "Deployment.Namespace", ingester.Namespace, "Deployment.Name", ingester.Name)
	err = r.client.Create(context.TODO(), ingester)
	if err != nil {
		reqLogger.Error(err, "Failed to create new Deployment", "Deployment.Namespace", ingester.Namespace, "Deployment.Name", ingester.Name)
		return err
	}
	ingesterService := r.serviceForIngester(instance)
	err = r.client.Create(context.TODO(), ingesterService)
	if err != nil {
		reqLogger.Error(err, "Failed to create new distributor service", "Deployment.Namespace", ingester.Namespace, "Deployment.Name", ingester.Name)
		return err
	}
	return nil
}

func (r *ReconcileCortex) createQuerier(instance *cortexv1alpha1.Cortex) error {
	deployment := &v1.Deployment{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: querier, Namespace: instance.Namespace}, deployment)
	if instance.Spec.Querier.Size == 0 || err == nil {
		return nil
	}
	reqLogger := log.WithValues("Request.Namespace", instance.Namespace, "Request.Name", instance.Name)
	querier := r.deploymentForQuerier(instance)
	reqLogger.Info("Creating a new Deployment", "Deployment.Namespace", querier.Namespace, "Deployment.Name", querier.Name)
	err = r.client.Create(context.TODO(), querier)
	if err != nil {
		reqLogger.Error(err, "Failed to create new Deployment", "Deployment.Namespace", querier.Namespace, "Deployment.Name", querier.Name)
		return err
	}
	querierService := r.serviceForQuerier(instance)
	err = r.client.Create(context.TODO(), querierService)
	if err != nil {
		reqLogger.Error(err, "Failed to create new distributor service", "Deployment.Namespace", querierService.Namespace, "Deployment.Name", querierService.Name)
		return err
	}
	return nil
}

func (r *ReconcileCortex) createQueryFrontEnd(instance *cortexv1alpha1.Cortex) error {
	deployment := &v1.Deployment{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: queryFrontEnd, Namespace: instance.Namespace}, deployment)
	if instance.Spec.QueryFrontEnd.Size == 0 || err == nil {
		return nil
	}
	reqLogger := log.WithValues("Request.Namespace", instance.Namespace, "Request.Name", instance.Name)
	queryFrontEnd := r.deploymentForQueryFrontEnd(instance)
	reqLogger.Info("Creating a new Deployment", "Deployment.Namespace", queryFrontEnd.Namespace, "Deployment.Name", queryFrontEnd.Name)
	err = r.client.Create(context.TODO(), queryFrontEnd)
	if err != nil {
		reqLogger.Error(err, "Failed to create new Deployment", "Deployment.Namespace", queryFrontEnd.Namespace, "Deployment.Name", queryFrontEnd.Name)
		return err
	}
	queryFrontEndService := r.serviceForQueryFrontEnd(instance)
	err = r.client.Create(context.TODO(), queryFrontEndService)
	if err != nil {
		reqLogger.Error(err, "Failed to create new distributor service", "Deployment.Namespace", queryFrontEndService.Namespace, "Deployment.Name", queryFrontEndService.Name)
		return err
	}
	return nil
}

func (r *ReconcileCortex) createTableManager(instance *cortexv1alpha1.Cortex) error {
	deployment := &v1.Deployment{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: tableManager, Namespace: instance.Namespace}, deployment)
	if instance.Spec.TableManager.Size == 0 || err == nil {
		return nil
	}
	reqLogger := log.WithValues("Request.Namespace", instance.Namespace, "Request.Name", instance.Name)
	tableManager := r.deploymentForTableManager(instance)
	reqLogger.Info("Creating a new Deployment", "Deployment.Namespace", tableManager.Namespace, "Deployment.Name", tableManager.Name)
	err = r.client.Create(context.TODO(), tableManager)
	if err != nil {
		reqLogger.Error(err, "Failed to create new Deployment", "Deployment.Namespace", tableManager.Namespace, "Deployment.Name", tableManager.Name)
		return err
	}
	tableManagerService := r.serviceForTableManager(instance)
	err = r.client.Create(context.TODO(), tableManagerService)
	if err != nil {
		reqLogger.Error(err, "Failed to create new distributor service", "Deployment.Namespace", tableManager.Namespace, "Deployment.Name", tableManager.Name)
		return err
	}
	return nil
}

func (r *ReconcileCortex) createS3(instance *cortexv1alpha1.Cortex) error {
	deployment := &v1.Deployment{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: s3, Namespace: instance.Namespace}, deployment)
	if instance.Spec.S3.Size == 0 || err == nil {
		return nil
	}
	reqLogger := log.WithValues("Request.Namespace", instance.Namespace, "Request.Name", instance.Name)
	s3 := r.deploymentForS3(instance)
	reqLogger.Info("Creating a new Deployment", "Deployment.Namespace", s3.Namespace, "Deployment.Name", s3.Name)
	err = r.client.Create(context.TODO(), s3)
	if err != nil {
		reqLogger.Error(err, "Failed to create new Deployment", "Deployment.Namespace", s3.Namespace, "Deployment.Name", s3.Name)
		return err
	}

	s3Service := r.serviceForS3(instance)
	err = r.client.Create(context.TODO(), s3Service)
	if err != nil {
		reqLogger.Error(err, "Failed to create new distributor service", "Deployment.Namespace", s3.Namespace, "Deployment.Name", s3.Name)
		return err
	}
	return nil
}

func (r *ReconcileCortex) createDynamoDB(instance *cortexv1alpha1.Cortex) error {
	deployment := &v1.Deployment{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: dynamoDB, Namespace: instance.Namespace}, deployment)
	if instance.Spec.DynamoDB.Size == 0 || err == nil {
		return nil
	}
	reqLogger := log.WithValues("Request.Namespace", instance.Namespace, "Request.Name", instance.Name)
	dynamo := r.deploymentForDynamoDB(instance)
	reqLogger.Info("Creating a new Deployment", "Deployment.Namespace", dynamo.Namespace, "Deployment.Name", dynamo.Name)
	err = r.client.Create(context.TODO(), dynamo)
	if err != nil {
		reqLogger.Error(err, "Failed to create new Deployment", "Deployment.Namespace", dynamo.Namespace, "Deployment.Name", dynamo.Name)
		return err
	}
	dynamoService := r.serviceForDynamoDb(instance)
	err = r.client.Create(context.TODO(), dynamoService)
	if err != nil {
		reqLogger.Error(err, "Failed to create new distributor service", "Deployment.Namespace", dynamo.Namespace, "Deployment.Name", dynamo.Name)
		return err
	}
	return nil
}

func (r *ReconcileCortex) createConfigs(instance *cortexv1alpha1.Cortex) error {
	deployment := &v1.Deployment{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: configs, Namespace: instance.Namespace}, deployment)
	if instance.Spec.Configs.Size == 0 || err == nil {
		return nil
	}
	reqLogger := log.WithValues("Request.Namespace", instance.Namespace, "Request.Name", instance.Name)
	configs := r.deploymentForConfigs(instance)
	reqLogger.Info("Creating a new Deployment", "Deployment.Namespace", configs.Namespace, "Deployment.Name", configs.Name)
	err = r.client.Create(context.TODO(), configs)
	if err != nil {
		reqLogger.Error(err, "Failed to create new Deployment", "Deployment.Namespace", configs.Namespace, "Deployment.Name", configs.Name)
		return err
	}
	configsService := r.serviceForConfigs(instance)
	err = r.client.Create(context.TODO(), configsService)
	if err != nil {
		reqLogger.Error(err, "Failed to create new distributor service", "Deployment.Namespace", configsService.Namespace, "Deployment.Name", configsService.Name)
		return err
	}
	return nil
}

func (r *ReconcileCortex) createConfigsDB(instance *cortexv1alpha1.Cortex) error {
	deployment := &v1.Deployment{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: configsDB, Namespace: instance.Namespace}, deployment)
	if instance.Spec.ConfigsDB.Size == 0 || err == nil {
		return nil
	}
	reqLogger := log.WithValues("Request.Namespace", instance.Namespace, "Request.Name", instance.Name)
	configsDb := r.deploymentForConfigsDb(instance)
	reqLogger.Info("Creating a new Deployment", "Deployment.Namespace", configsDb.Namespace, "Deployment.Name", configsDb.Name)
	err = r.client.Create(context.TODO(), configsDb)
	if err != nil {
		reqLogger.Error(err, "Failed to create new Deployment", "Deployment.Namespace", configsDb.Namespace, "Deployment.Name", configsDb.Name)
		return err
	}
	configsDbService := r.serviceForConfigsDb(instance)
	err = r.client.Create(context.TODO(), configsDbService)
	if err != nil {
		reqLogger.Error(err, "Failed to create new distributor service", "Deployment.Namespace", configsDbService.Namespace, "Deployment.Name", configsDbService.Name)
		return err
	}
	return nil
}

func (r *ReconcileCortex) createNginx(instance *cortexv1alpha1.Cortex) error {
	deployment := &v1.Deployment{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: nginx, Namespace: instance.Namespace}, deployment)
	if instance.Spec.Nginx.Size == 0 || err == nil {
		return nil
	}

	reqLogger := log.WithValues("Request.Namespace", instance.Namespace, "Request.Name", instance.Name)
	nginxConfigMap := r.nginxConfigMap(instance)
	reqLogger.Info("Creating a new Deployment", "Deployment.Namespace", nginxConfigMap.Namespace, "Deployment.Name", nginxConfigMap.Name)
	err = r.client.Create(context.TODO(), nginxConfigMap)
	if err != nil {
		reqLogger.Error(err, "Failed to create new configmap", "Deployment.Namespace", nginxConfigMap.Namespace, "Deployment.Name", nginxConfigMap.Name)
		return err
	}

	nginx := r.deploymentForNginx(instance)
	reqLogger.Info("Creating a new Deployment", "Deployment.Namespace", nginx.Namespace, "Deployment.Name", nginx.Name)
	err = r.client.Create(context.TODO(), nginx)
	if err != nil {
		reqLogger.Error(err, "Failed to create new Deployment", "Deployment.Namespace", nginx.Namespace, "Deployment.Name", nginx.Name)
		return err
	}

	nginxService := r.serviceForNginx(instance)
	err = r.client.Create(context.TODO(), nginxService)
	if err != nil {
		reqLogger.Error(err, "Failed to create new distributor service", "Deployment.Namespace", nginx.Namespace, "Deployment.Name", nginx.Name)
		return err
	}
	return nil
}

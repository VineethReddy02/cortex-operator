package cortex

import (
	cortexv1alpha1 "github.com/VineethReddy02/goModules/cortex-operator/pkg/apis/cortex/v1alpha1"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// deploymentForIngester returns a ingester Deployment object
func (r *ReconcileCortex) deploymentForIngester(m *cortexv1alpha1.Cortex) *v1.Deployment {
	ls := labelsForIngester()
	replicas := m.Spec.Ingester.Size

	dep := &v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ingester,
			Namespace: m.Namespace,
		},
		Spec: v1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: ls,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: ls,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image:   "quay.io/cortexproject/cortex:master-f5e3dd77",
						Name:    ingester,
						Args:    []string{"-target=ingester", "-ingester.join-after=30s", "-ingester.claim-on-rollout=true", "-consul.hostname=consul.default.svc.cluster.local:8500", "-s3.url=s3://abc:123@s3.default.svc.cluster.local:4569", "-dynamodb.original-table-name=cortex", "-dynamodb.url=dynamodb://user:pass@dynamodb.default.svc.cluster.local:8000", "-dynamodb.periodic-table.prefix=cortex_weekly_", "-dynamodb.periodic-table.from=2017-01-06", "-dynamodb.daily-buckets-from=2017-01-10", "-dynamodb.base64-buckets-from=2017-01-17", "-dynamodb.v4-schema-from=2017-02-05", "-dynamodb.v5-schema-from=2017-02-22", "-dynamodb.v6-schema-from=2017-03-19", "-dynamodb.chunk-table.from=2017-04-17", "-memcached.hostname=memcached.default.svc.cluster.local", "-memcached.timeout=100ms", "-memcached.service=memcached"},
						Ports: []corev1.ContainerPort{{
							ContainerPort: 80,
							Name:          "ingester",
						}},
					}},
				},
			},
		},
	}

	controllerutil.SetControllerReference(m, dep, r.scheme)
	return dep
}

// deploymentForIngester returns a ingester Deployment object
func (r *ReconcileCortex) deploymentForDistributor(m *cortexv1alpha1.Cortex) *v1.Deployment {
	ls := labelsForDistributor()
	replicas := m.Spec.Distributor.Size

	dep := &v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      distributor,
			Namespace: m.Namespace,
		},
		Spec: v1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: ls,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: ls,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image:   "quay.io/cortexproject/cortex:master-f5e3dd77",
						Name:    distributor,
						Args:    []string{"-target=distributor", "-log.level=debug", "-server.http-listen-port=80", "-consul.hostname=consul.default.svc.cluster.local:8500", "-distributor.replication-factor=1"},
						Ports: []corev1.ContainerPort{{
							ContainerPort: 80,
							Name:          "distributor",
						}},
					}},
				},
			},
		},
	}

	controllerutil.SetControllerReference(m, dep, r.scheme)
	return dep
}

// deploymentForIngester returns a ingester Deployment object
func (r *ReconcileCortex) deploymentForQuerier(m *cortexv1alpha1.Cortex) *v1.Deployment {
	ls := labelsForQuerier()
	replicas := m.Spec.Querier.Size

	dep := &v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      querier,
			Namespace: m.Namespace,
		},
		Spec: v1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: ls,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: ls,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image:   "quay.io/cortexproject/cortex:master-f5e3dd77",
						Name:    "querier",
						Args:    []string{"-target=querier", "-server.http-listen-port=80", "-consul.hostname=consul.default.svc.cluster.local:8500", "-s3.url=s3://abc:123@s3.default.svc.cluster.local:4569", "-querier.frontend-address=query-frontend.default.svc.cluster.local:9095", "-dynamodb.original-table-name=cortex", "-dynamodb.url=dynamodb://user:pass@dynamodb.default.svc.cluster.local:8000", "-dynamodb.periodic-table.prefix=cortex_weekly_", "-dynamodb.periodic-table.from=2017-01-06", "-dynamodb.daily-buckets-from=2017-01-10", "-dynamodb.base64-buckets-from=2017-01-17", "-dynamodb.v4-schema-from=2017-02-05", "-dynamodb.v5-schema-from=2017-02-22", "-dynamodb.v6-schema-from=2017-03-19", "-dynamodb.chunk-table.from=2017-04-17", "-memcached.hostname=memcached.default.svc.cluster.local", "-memcached.timeout=100ms", "-memcached.service=memcached", "-distributor.replication-factor=1"},
						Ports: []corev1.ContainerPort{{
							ContainerPort: 80,
							Name:          "querier",
						}},
					}},
				},
			},
		},
	}

	controllerutil.SetControllerReference(m, dep, r.scheme)
	return dep
}

// deploymentForIngester returns a ingester Deployment object
func (r *ReconcileCortex) deploymentForConsul(m *cortexv1alpha1.Cortex) *v1.Deployment {
	ls := labelsForConsul()
	replicas := m.Spec.Consul.Size

	dep := &v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      consul,
			Namespace: m.Namespace,
		},
		Spec: v1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: ls,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: ls,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image:   "consul:0.7.1",
						Name:    consul,
						Args:    []string{"agent", "-ui", "-server", "-client=0.0.0.0", "-bootstrap"},
						Ports: []corev1.ContainerPort{{
							ContainerPort: 8300,
							Name:          "server-noscrape",
						},
							{
								ContainerPort: 8301,
								Name:          "serf-noscrape",
							},
							{
								ContainerPort: 8400,
								Name:          "client-noscrape",
							},
							{
								ContainerPort: 8500,
								Name:          "http-noscrape",
							},
						},
						Env: []corev1.EnvVar{{
							Name:      "CHECKPOINT_DISABLE",
							Value:     "1",
						}},
					}},
				},
			},
		},
	}

	controllerutil.SetControllerReference(m, dep, r.scheme)
	return dep
}

func (r *ReconcileCortex) deploymentForNginx(m *cortexv1alpha1.Cortex) *v1.Deployment {
	ls := labelsForNginx()
	replicas := m.Spec.Nginx.Size

	dep := &v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      nginx,
			Namespace: m.Namespace,
		},
		Spec: v1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: ls,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: ls,
					Annotations: map[string]string{"prometheus.io.scrape":"false"},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image:   "nginx",
						Name:    nginx,
						Ports: []corev1.ContainerPort{{
							ContainerPort: 80,
							Name:          "http",
						},
						},
						VolumeMounts: []corev1.VolumeMount{{
							Name: "config-volume",
							MountPath: "/etc/nginx",
						}},
					}},
					Volumes: []corev1.Volume{{
						Name: "config-volume",
						VolumeSource: corev1.VolumeSource{
							ConfigMap:             &corev1.ConfigMapVolumeSource{
								LocalObjectReference: corev1.LocalObjectReference{Name:"nginx"},
							},
						},
					}},
				},
			},
		},
	}

	controllerutil.SetControllerReference(m, dep, r.scheme)
	return dep
}

func (r *ReconcileCortex) deploymentForS3(m *cortexv1alpha1.Cortex) *v1.Deployment {
	ls := labelsForS3()
	replicas := m.Spec.S3.Size

	dep := &v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      s3,
			Namespace: m.Namespace,
		},
		Spec: v1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: ls,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: ls,
					Annotations: map[string]string{"prometheus.io.scrape":"false"},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image:   "lphoward/fake-s3",
						Name:    s3,
						Ports: []corev1.ContainerPort{{
							ContainerPort: 4569,
						},
						},
					}},
				},
			},
		},
	}

	controllerutil.SetControllerReference(m, dep, r.scheme)
	return dep
}

func (r *ReconcileCortex) deploymentForDynamoDB(m *cortexv1alpha1.Cortex) *v1.Deployment {
	ls := labelsForDynamo()
	replicas := m.Spec.DynamoDB.Size

	dep := &v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      dynamoDB,
			Namespace: m.Namespace,
		},
		Spec: v1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: ls,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: ls,
					Annotations: map[string]string{"prometheus.io.scrape":"false"},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image:   "amazon/dynamodb-local:1.11.477",
						Name:    dynamoDB,
						Args:    []string{"-jar","DynamoDBLocal.jar","-inMemory","-sharedDb"},
						Ports: []corev1.ContainerPort{{
							ContainerPort: 8000,
						},
						},
					}},
				},
			},
		},
	}

	controllerutil.SetControllerReference(m, dep, r.scheme)
	return dep
}

func (r *ReconcileCortex) deploymentForTableManager(m *cortexv1alpha1.Cortex) *v1.Deployment {
	ls := labelsForTableManager()
	replicas := m.Spec.TableManager.Size

	dep := &v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      tableManager,
			Namespace: m.Namespace,
		},
		Spec: v1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: ls,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: ls,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image:   "quay.io/cortexproject/cortex:master-f5e3dd77",
						Name:    tableManager,
						Args:    []string{"-target=table-manager", "-server.http-listen-port=80", "-dynamodb.original-table-name=cortex", "-dynamodb.url=dynamodb://user:pass@dynamodb.default.svc.cluster.local:8000", "-dynamodb.periodic-table.prefix=cortex_weekly_","-dynamodb.periodic-table.from=2017-01-06","-dynamodb.chunk-table.from=2017-04-17"},
						Ports: []corev1.ContainerPort{{
							ContainerPort: 80,
						},
						},
					}},
				},
			},
		},
	}

	controllerutil.SetControllerReference(m, dep, r.scheme)
	return dep
}

func (r *ReconcileCortex) deploymentForConfigs(m *cortexv1alpha1.Cortex) *v1.Deployment {
	ls := labelsForConfigs()
	replicas := m.Spec.Configs.Size

	dep := &v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      configs,
			Namespace: m.Namespace,
		},
		Spec: v1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: ls,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: ls,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image:   "quay.io/cortexproject/cortex:master-f5e3dd77",
						Name:    configs,
						Args:    []string{"-target=configs", "-server.http-listen-port=80", "-database.uri=postgres://postgres@configs-db.default.svc.cluster.local/configs?sslmode=disable", "-database.migrations=/migrations"},
						Ports: []corev1.ContainerPort{{
							ContainerPort: 80,
						},
						},
					}},
				},
			},
		},
	}

	controllerutil.SetControllerReference(m, dep, r.scheme)
	return dep
}

func (r *ReconcileCortex) deploymentForConfigsDb(m *cortexv1alpha1.Cortex) *v1.Deployment {
	ls := labelsForConfigsDB()
	replicas := m.Spec.ConfigsDB.Size

	dep := &v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      configsDB,
			Namespace: m.Namespace,
		},
		Spec: v1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: ls,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: ls,
					Annotations: map[string]string{"prometheus.io.scrape":"false"},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image:   "postgres:9.6",
						Name:    configsDB,
						Ports: []corev1.ContainerPort{{
							ContainerPort: 5432,
						},
						},
						Env:  []corev1.EnvVar{{
							Name:      "POSTGRES_DB",
							Value:     "configs",
						}},
					}},
				},
			},
		},
	}

	controllerutil.SetControllerReference(m, dep, r.scheme)
	return dep
}

func (r *ReconcileCortex) deploymentForQueryFrontEnd(m *cortexv1alpha1.Cortex) *v1.Deployment {
	ls := labelsForQueryFrontEnd()
	replicas := m.Spec.QueryFrontEnd.Size

	dep := &v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      queryFrontEnd,
			Namespace: m.Namespace,
		},
		Spec: v1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: ls,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: ls,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image:   "quay.io/cortexproject/cortex:master-f5e3dd77",
						Name:    queryFrontEnd,
						Args:    []string{"-target=query-frontend","-log.level=debug","-server.http-listen-port=80","-server.grpc-listen-port=9095"},
						Ports: []corev1.ContainerPort{{
							ContainerPort: 9095,
							Name:          "grpc",
						},
							{
								ContainerPort: 80,
								Name:          "http",
							},
						},
					}},
				},
			},
		},
	}

	controllerutil.SetControllerReference(m, dep, r.scheme)
	return dep
}

func labelsForIngester() map[string]string {
	return map[string]string{"name": ingester}
}

func labelsForDistributor() map[string]string {
	return map[string]string{"name": distributor}
}

func labelsForQuerier() map[string]string {
	return map[string]string{"name": querier}
}

func labelsForConsul() map[string]string {
	return map[string]string{"name": consul}
}

func labelsForNginx() map[string]string {
	return map[string]string{"name": nginx}
}

func labelsForConfigsDB() map[string]string {
	return map[string]string{"name": configsDB}
}

func labelsForConfigs() map[string]string {
	return map[string]string{"name": configs}
}

func labelsForS3() map[string]string {
	return map[string]string{"name": s3}
}

func labelsForDynamo() map[string]string {
	return map[string]string{"name": dynamoDB}
}

func labelsForTableManager() map[string]string {
	return map[string]string{"name": tableManager}
}

func labelsForQueryFrontEnd() map[string]string {
	return map[string]string{"name": queryFrontEnd}
}
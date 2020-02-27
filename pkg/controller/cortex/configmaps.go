package cortex

import (
	cortexv1alpha1 "github.com/VineethReddy02/goModules/cortex-operator/pkg/apis/cortex/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func (r *ReconcileCortex) nginxConfigMap(m *cortexv1alpha1.Cortex) *corev1.ConfigMap {
	configMap := &corev1.ConfigMap{
		TypeMeta:   metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:                       "nginx",
			Namespace:                 m.Namespace,
		},
		Data:       map[string]string{"nginx.conf":`
			worker_processes  5;  ## Default: 1
			error_log  /dev/stderr;
			pid        /tmp/nginx.pid;
			worker_rlimit_nofile 8192;

			events {
			worker_connections  4096;  ## Default: 1024
		}

			http {
			default_type application/octet-stream;
			log_format   main '$remote_addr - $remote_user [$time_local]  $status '
			'"$request" $body_bytes_sent "$http_referer" '
			'"$http_user_agent" "$http_x_forwarded_for"';
			access_log   /dev/stderr  main;
			sendfile     on;
			tcp_nopush   on;
			resolver kube-dns.kube-system.svc.cluster.local;

			server { # simple reverse-proxy
			listen 80;
			proxy_set_header X-Scope-OrgID 0;

			location = /api/prom/push {
			proxy_pass      http://distributor.default.svc.cluster.local$request_uri;
		}

			location = /ring {
			proxy_pass      http://distributor.default.svc.cluster.local$request_uri;
		}
			location = /all_user_stats {
			proxy_pass      http://distributor.default.svc.cluster.local$request_uri;
		}

			location ~ /api/prom/.* {
			proxy_pass      http://query-frontend.default.svc.cluster.local$request_uri;
		}
		}
		}`},
	}

	controllerutil.SetControllerReference(m, configMap, r.scheme)
	return configMap
}

# Cortex Operator

This is the initial work done to ease Cortex deployment in Kubernetes environment.

Now you just need the operator, cortex custom resource definition and cortex custom resource.

##### Create Cortex Custom Resource Definition 

```kubectl create -f deploy/crds/cortex.k8s.com_cortexes_crd.yaml```

##### Create Cortex Operator

```kubectl create -f deploy/operator.yaml``` 

Configure the Cortex custom resource as per your requirement. At this point of time you can configure required cortex services and number of replicas for each Cortex service.

```cassandraql
apiVersion: cortex.k8s.com/v1alpha1
kind: Cortex
metadata:
  name: example-cortex
spec:
  consul:
    size: 1
  distributor:
    size: 1
  ingester:
    size: 1
  querier:
    size: 1
  queryFrontEnd:
    size: 1
  nginx:
    size: 1
  tableManager:
    size: 1
  s3:
    size: 1
  dynamoDB:
    size: 1
```

##### Create above configured Cortex custom resource

```kubectl create -f deploy/crds/cortex.k8s.com_v1alpha1_cortex_cr.yaml```

Now Cortex is successfully deployed:

```cassandraql
➜ kubectl get cortex
NAME             AGE
cortex           69s
```

```cassandraql
➜ kubectl get deployments
NAME              READY   UP-TO-DATE   AVAILABLE   AGE
consul            1/1     1            1           2m8s
cortex-operator   1/1     1            1           2m20s
distributor       1/1     1            1           2m8s
dynamodb          1/1     1            1           2m7s
ingester          1/1     1            1           2m8s
nginx             1/1     1            1           2m8s
querier           1/1     1            1           2m7s
query-frontend    1/1     1            1           2m8s
s3                1/1     1            1           2m7s
table-manager     1/1     1            1           2m7s

```

Also you can check the ring status using this endpoint:
```cassandraql
http://<NODE_IP>:30080/ring
```

##### TODO's:

1. Configurable storage backend using the custom resource.
2. Configurable ring i.e consul/etcd/gossip using the custom resource.
3. Make Cortex custom resource more configurable i.e specific to each Cortex service.
4. Add more operational logic for easy maintainence in case of downtime of specific service.
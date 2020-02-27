package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CortexSpec defines the desired state of Cortex
type CortexSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Ingester      IngesterSpec      `json:"ingester,omitempty"`
	Nginx         NginxSpec         `json:"nginx,omitempty"`
	Consul        ConsulSpec        `json:"consul,omitempty"`
	Distributor   DistributorSpec   `json:"distributor,omitempty"`
	Querier       QuerierSpec       `json:"querier,omitempty"`
	QueryFrontEnd QueryFrontEndSpec `json:"queryFrontEnd,omitempty"`
	TableManager  TableManagerSpec  `json:"tableManager,omitempty"`
	Configs       ConfigsSpec       `json:"configs,omitempty"`
	ConfigsDB     ConfigsDbSpec     `json:"configsDB,omitempty"`
	AlertManager  AlertManagerSpec  `json:"alertManager,omitempty"`
	Ruler         RulerSpec         `json:"ruler,omitempty"`
	Memcached     MemcachedSpec     `json:"memcached,omitempty"`
	S3            S3Spec            `json:"s3,omitempty"`
	DynamoDB      DynamoDBSpec      `json:"dynamodb,omitempty"`
}

// CortexStatus defines the observed state of Cortex
type CortexStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Nodes []string `json:"nodes"`
}

type IngesterSpec struct {
	Size int32 `json:"size"`
}

type NginxSpec struct {
	Size int32 `json:"size"`
}

type ConsulSpec struct {
	Size int32 `json:"size"`
}

type DistributorSpec struct {
	Size int32 `json:"size"`
}

type QuerierSpec struct {
	Size int32 `json:"size"`
}

type QueryFrontEndSpec struct {
	Size int32 `json:"size"`
}

type TableManagerSpec struct {
	Size int32 `json:"size"`
}

type ConfigsSpec struct {
	Size int32 `json:"size"`
}

type ConfigsDbSpec struct {
	Size int32 `json:"size"`
}

type AlertManagerSpec struct {
	Size int32 `json:"size"`
}

type RulerSpec struct {
	Size int32 `json:"size"`
}

type MemcachedSpec struct {
	Size int32 `json:"size"`
}

type S3Spec struct {
	Size int32 `json:"size"`
}

type DynamoDBSpec struct {
	Size int32 `json:"size"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Cortex is the Schema for the cortexes API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=cortexes,scope=Namespaced
type Cortex struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CortexSpec   `json:"spec,omitempty"`
	Status CortexStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CortexList contains a list of Cortex
type CortexList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cortex `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Cortex{}, &CortexList{})
}

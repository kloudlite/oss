package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ct "operators.kloudlite.io/apis/common-types"
	rApi "operators.kloudlite.io/lib/operator"
)

type RedpandaAdmin struct {
	SecretRef ct.SecretRef `json:"secretRef"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	KafkaTopics   []string      `json:"kafkaTopics,omitempty"`
	RedpandaAdmin RedpandaAdmin `json:"redpandaAdmin"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:JSONPath=".status.isReady",name=Ready,type=boolean
// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name=Age,type=date

// Cluster is the Schema for the clusters API
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterSpec `json:"spec,omitempty"`
	Status rApi.Status `json:"status,omitempty"`
}

func (c *Cluster) GetStatus() *rApi.Status {
	return &c.Status
}

func (c *Cluster) GetEnsuredLabels() map[string]string {
	return map[string]string{}
}

func (in *Cluster) GetEnsuredAnnotations() map[string]string {
	return map[string]string{}
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Cluster
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}

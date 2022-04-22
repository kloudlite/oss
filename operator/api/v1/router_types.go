package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type routes struct {
	Path string `json:"path"`
	App  string `json:"app"`
	Port uint16 `json:"port"`
}

// RouterSpec defines the desired state of Router
type RouterSpec struct {
	Domains []string `json:"domains"`
	Routes  []routes `json:"routes"`
}

// RouterStatus defines the observed state of Router
type RouterStatus struct {
	IngressCheck Recon              `json:"ingress_check"`
	Conditions   []metav1.Condition `json:"conditions"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Router is the Schema for the routers API
type Router struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RouterSpec   `json:"spec,omitempty"`
	Status RouterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RouterList contains a list of Router
type RouterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Router `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Router{}, &RouterList{})
}

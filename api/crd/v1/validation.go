package v1

//Validation CRD objects structure

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Validation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ValidationSpec `json:"spec"`
	Status            StatusInfo     `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ValidationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Validation `json:"items"`
}

// +k8s:openapi-gen=true
type ValidationSpec struct {
	Name       string `json:"name,omitempty"`
	Namespace  string `json:"namespace,omitempty"`
	Resource   string `json:"resource,omitempty"`
	Validation []struct {
		JsonPath string `json:"jsonPath"`
		Value    string `json:"value"`
	} `json:"validation"`
}

type StatusInfo struct {
	LastCheck   string `json:"lastCheck"`
	LastChanged string `json:"lastChange"`
	State       string `json:"state"`
}

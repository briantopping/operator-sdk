package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// FreeIPAServiceSpec defines the desired state of FreeIPAService
// +k8s:openapi-gen=true
type FreeIPAServiceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// FreeIPAServiceStatus defines the observed state of FreeIPAService
// +k8s:openapi-gen=true
type FreeIPAServiceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FreeIPAService is the Schema for the freeipaservices API
// +k8s:openapi-gen=true
type FreeIPAService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FreeIPAServiceSpec   `json:"spec,omitempty"`
	Status FreeIPAServiceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FreeIPAServiceList contains a list of FreeIPAService
type FreeIPAServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FreeIPAService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FreeIPAService{}, &FreeIPAServiceList{})
}

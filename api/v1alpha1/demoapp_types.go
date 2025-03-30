package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DemoAppSpec defines the desired state of DemoApp
type DemoAppSpec struct {
	// Message is a simple message that will be displayed
	Message string `json:"message"`
}

// DemoAppStatus defines the observed state of DemoApp
type DemoAppStatus struct {
	// Phase represents the current phase of the DemoApp
	Phase string `json:"phase,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.phase"
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// DemoApp is the Schema for the demoapps API
type DemoApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DemoAppSpec   `json:"spec,omitempty"`
	Status DemoAppStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DemoAppList contains a list of DemoApp
type DemoAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DemoApp `json:"items"`
} 
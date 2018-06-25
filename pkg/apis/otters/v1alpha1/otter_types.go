package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!
// Created by "kubebuilder create resource" for you to implement the Otter resource schema definition
// as a go struct.
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
// Important: Run "kubebuilder generate" to regenerate code after modifying this file

// OtterSpec defines the desired state of Otter
type OtterSpec struct {
	Replicas int32 `json:"replicas,omitempty"`

	// image is the container image to run.  Image must have a tag.
    // +kubebuilder:validation:Pattern=.+:.+
	Image string `json:"image,omitempty"`
}

// OtterStatus defines the observed state of Otter
type OtterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "kubebuilder generate" to regenerate code after modifying this file
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Otter creates a new Deployment running multiple replicas of a single 
// container named after a species of otter
// +k8s:openapi-gen=true
// +kubebuilder:resource:path=otters
type Otter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// spec contains the desired behavior of the Otter
	Spec   OtterSpec   `json:"spec,omitempty"`

	// status contains the last observed state of the Otter
	Status OtterStatus `json:"status,omitempty"`
}

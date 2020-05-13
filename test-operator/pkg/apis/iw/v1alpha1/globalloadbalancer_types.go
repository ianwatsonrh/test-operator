package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GlobalLoadBalancerSpec defines the desired state of GlobalLoadBalancer
type GlobalLoadBalancerSpec struct {
	Name string `json:"name"`
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// GlobalLoadBalancerStatus defines the observed state of GlobalLoadBalancer
type GlobalLoadBalancerStatus struct {
        Status string `json:"status"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GlobalLoadBalancer is the Schema for the globalloadbalancers API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=globalloadbalancers,scope=Namespaced
type GlobalLoadBalancer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlobalLoadBalancerSpec   `json:"spec,omitempty"`
	Status GlobalLoadBalancerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GlobalLoadBalancerList contains a list of GlobalLoadBalancer
type GlobalLoadBalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalLoadBalancer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GlobalLoadBalancer{}, &GlobalLoadBalancerList{})
}

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

// GroupVersion for the Rule API
var GroupVersion = schema.GroupVersion{Group: "kwaf.io", Version: "v1alpha1"}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
	Codecs        = serializer.NewCodecFactory(runtime.NewScheme())
)

// RuleSpec defines the desired state of a WAF rule
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type RuleSpec struct {
	ID      string `json:"id"`
	Pattern string `json:"pattern"`
	Action  string `json:"action"`
}

// Rule is the Schema for the rules API.
type Rule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec RuleSpec `json:"spec,omitempty"`
}

// RuleList contains a list of Rule
type RuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Rule `json:"items"`
}

func (r *Rule) DeepCopyObject() runtime.Object {
	out := new(Rule)
	*out = *r
	return out
}

func (rl *RuleList) DeepCopyObject() runtime.Object {
	out := new(RuleList)
	*out = *rl
	return out
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(GroupVersion, &Rule{}, &RuleList{})
	metav1.AddToGroupVersion(scheme, GroupVersion)
	return nil
}

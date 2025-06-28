package v1alpha1

// RuleSpec defines the desired state of a WAF rule
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type RuleSpec struct {
	ID      string `json:"id"`
	Pattern string `json:"pattern"`
	Action  string `json:"action"`
}

// Rule is the Schema for the rules API.
// This simplified struct does not include full Kubernetes metadata.
type Rule struct {
	Spec RuleSpec `json:"spec"`
}

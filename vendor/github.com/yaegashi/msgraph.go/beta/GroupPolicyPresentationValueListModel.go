// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// GroupPolicyPresentationValueList The entity represents a collection of name/value pairs of a list box presentation on a policy definition.
type GroupPolicyPresentationValueList struct {
	// GroupPolicyPresentationValue is the base model of GroupPolicyPresentationValueList
	GroupPolicyPresentationValue
	// Values A list of pairs for the associated presentation.
	Values []KeyValuePair `json:"values,omitempty"`
}
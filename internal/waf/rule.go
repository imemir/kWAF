package waf

import "regexp"

// Rule defines a single pattern based security rule.
type Rule struct {
	ID      string
	Pattern *regexp.Regexp
	Action  string // e.g. "block" or "allow"
}

// NewRule compiles the given pattern and returns a Rule instance.
func NewRule(id, pattern, action string) (*Rule, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return &Rule{ID: id, Pattern: re, Action: action}, nil
}

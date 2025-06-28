package waf

import (
	"regexp"
)

// Rule defines a single pattern based security rule.
type Rule struct {
	ID string
	// Pattern stores the raw pattern string as provided by the user.
	Pattern string
	// Type determines how the pattern should be evaluated. Supported values
	// are "regex" (default) and "substring".
	Type   string
	Action string // e.g. "block" or "allow"

	re *regexp.Regexp
}

// NewRule compiles the given pattern and returns a Rule instance.
func NewRule(id, pattern, action string) (*Rule, error) {
	return NewRuleWithType(id, pattern, action, "regex")
}

// NewRuleWithType creates a rule with an explicit pattern type.
func NewRuleWithType(id, pattern, action, typ string) (*Rule, error) {
	r := &Rule{ID: id, Pattern: pattern, Action: action, Type: typ}
	if typ == "regex" || typ == "" {
		re, err := regexp.Compile(pattern)
		if err != nil {
			return nil, err
		}
		r.re = re
		r.Type = "regex"
	}
	return r, nil
}

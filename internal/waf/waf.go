package waf

import "strings"

// Engine represents the core firewall engine
// This minimal implementation evaluates inputs against a set of rules.
type Engine struct {
	rules map[string]*Rule
}

// NewEngine creates a new firewall engine
func NewEngine() *Engine {
	return &Engine{rules: make(map[string]*Rule)}
}

// UpsertRule adds or updates a compiled rule in the engine.
func (e *Engine) UpsertRule(r *Rule) {
	if e.rules == nil {
		e.rules = make(map[string]*Rule)
	}
	e.rules[r.ID] = r
}

func (e *Engine) DeleteRule(id string) {
	delete(e.rules, id)
}

// Evaluate checks the input against the loaded rules and returns true if the
// input should be blocked.
func (e *Engine) Evaluate(input string) bool {
	for _, r := range e.rules {
		var match bool
		switch r.Type {
		case "substring":
			match = strings.Contains(input, r.Pattern)
		default: // regex
			if r.re != nil {
				match = r.re.MatchString(input)
			}
		}

		if match {
			switch r.Action {
			case "allow":
				return false
			case "block":
				return true
			}
		}
	}
	return false
}

package waf

import "testing"

func TestEngineEvaluate(t *testing.T) {
	e := NewEngine()
	r, err := NewRule("test", `(?i)drop table`, "block")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	e.UpsertRule(r)

	if !e.Evaluate("user tries to DROP TABLE users") {
		t.Errorf("expected input to be blocked")
	}

	if e.Evaluate("safe input") {
		t.Errorf("expected input to be allowed")
	}
}

func TestEngineEvaluate_Substring(t *testing.T) {
	e := NewEngine()
	r, err := NewRuleWithType("substr", "admin", "block", "substring")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	e.UpsertRule(r)
	if !e.Evaluate("account=admin") {
		t.Errorf("expected substring rule to block input")
	}
}

func TestEngineEvaluate_Allow(t *testing.T) {
	e := NewEngine()
	allow, err := NewRuleWithType("allow", "trusted", "allow", "substring")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	block, err := NewRule("sql", `(?i)drop table`, "block")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	e.UpsertRule(allow)
	e.UpsertRule(block)

	if e.Evaluate("trusted user tries to DROP TABLE foo") {
		t.Errorf("expected allow rule to override block")
	}

	if !e.Evaluate("DROP TABLE bar") {
		t.Errorf("expected block rule to block input")
	}
}

package waf

import "testing"

func TestEngineEvaluate(t *testing.T) {
	e := NewEngine()
	r, err := NewRule("test", `(?i)drop table`, "block")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	e.AddRule(r)

	if !e.Evaluate("user tries to DROP TABLE users") {
		t.Errorf("expected input to be blocked")
	}

	if e.Evaluate("safe input") {
		t.Errorf("expected input to be allowed")
	}
}

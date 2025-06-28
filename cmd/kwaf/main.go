package main

import (
	"fmt"
	"os"

	"kwaf/internal/waf"
)

// simple CLI demonstrating rule evaluation
func main() {
	fmt.Println("kWAF starting...")

	// create engine and add a sample rule
	eng := waf.NewEngine()
	rule, err := waf.NewRule("sql-injection", `(?i)drop table`, "block")
	if err != nil {
		fmt.Println("invalid rule pattern:", err)
		return
	}
	eng.AddRule(rule)

	if len(os.Args) < 2 {
		fmt.Println("usage: kwaf <input>")
		return
	}
	input := os.Args[1]
	if eng.Evaluate(input) {
		fmt.Println("Blocked input:", input)
	} else {
		fmt.Println("Allowed input:", input)
	}
}

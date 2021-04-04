package main

import (
	"regexp"
	"testing"
)

func TestTestSetup(t *testing.T) {
	text := "Hello World"
	want := regexp.MustCompile(`\b` + "Hello World" + `\b`)
	if !want.MatchString(text) {
		t.Fatalf(`text = %q, want match for %v`, text, want)
	}
}

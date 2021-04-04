package main

import (
	"testing"
)

func TestCanTest(t *testing.T) {
	var result int = 1
	if result != 1 {
		t.Error("Failed to test.")
	}
}

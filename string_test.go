package main

import (
	"testing"
)

func TestBalancedStringSplit(t *testing.T) {
	s := "RLRRLLRLRL"
	expected := 4
	actual := balancedStringSplit(s)
	if expected != actual {
		t.Fatalf("BalancedStringSplit failed. Expected: %v, got: %v", expected, actual)
	}
}

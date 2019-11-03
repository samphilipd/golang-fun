package main

import (
	"fmt"
	"testing"
)

func TestDefangIPaddr(t *testing.T) {
	input := "1.1.1.1"
	expected := "1[.]1[.]1[.]1"
	actual := defangIPaddr(input)
	if actual != expected {
		t.Fatalf("defangIPaddr failed: expected: %v, got: %v", expected, actual)
	} else {
		fmt.Println("TestDefangIPaddr pass")
	}
}
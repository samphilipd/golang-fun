package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	input := []int{3, 2, 3}
	expected := 3
	actual := MajorityElement(input)
	if actual != expected {
		t.Fatalf("MajorityElement failed. Expected: %v, got: %v", expected, actual)
	}

	input2 := []int{2, 2, 1, 1, 1, 2, 2}
	expected2 := 2
	actual2 := MajorityElement(input2)
	if actual2 != expected2 {
		t.Fatalf("MajorityElement failed. Expected: %v, got: %v", expected2, actual2)
	} else {
		fmt.Println("TestMajorityElement pass")
	}
}

func TestMajorityElements(t *testing.T) {
	input := []int{3, 2, 3}
	expected := []int{3}
	actual := MajorityElements(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("MajorityElement failed. Expected: %v, got: %v", expected, actual)
	}

	input2 := []int{1, 1, 1, 3, 3, 2, 2, 2}
	expected2 := []int{1, 2}
	actual2 := MajorityElements(input2)
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("MajorityElement failed. Expected: %v, got: %v", expected2, actual2)
	} else {
		fmt.Println("TestMajorityElement pass")
	}
}

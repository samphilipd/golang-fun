package main

import (
	"fmt"
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

func TestMajorityElementMoores(t *testing.T) {
	input := []int{3, 2, 3}
	expected := 3
	actual := MajorityElementMoores(input)
	if actual != expected {
		t.Fatalf("MajorityElementMoores failed. Expected: %v, got: %v", expected, actual)
	}

	input2 := []int{2, 2, 1, 1, 1, 2, 2}
	expected2 := 2
	actual2 := MajorityElementMoores(input2)
	if actual2 != expected2 {
		t.Fatalf("MajorityElementMoores failed. Expected: %v, got: %v", expected2, actual2)
	} else {
		fmt.Println("TestMajorityElementMoores pass")
	}
}

func TestMajorityElements(t *testing.T) {
	input := []int{3, 2, 3}
	expected := []int{3}
	actual := MajorityElements(input)
	if !ArraysEqual(actual, expected) {
		t.Fatalf("TestMajorityElements failed. Expected: %v, got: %v", expected, actual)
	}

	input2 := []int{1, 1, 1, 3, 3, 2, 2, 2}
	expected2 := []int{1, 2}
	actual2 := MajorityElements(input2)
	if !ArraysEqual(actual, expected) {
		t.Fatalf("TestMajorityElements failed. Expected: %v, got: %v", expected2, actual2)
	}

	input3 := []int{2, 2}
	expected3 := []int{2}
	actual3 := MajorityElements(input3)
	if !ArraysEqual(actual3, expected3) {
		t.Fatalf("TestMajorityElements failed. Expected: %v, got: %v", expected3, actual3)
	}

	input4 := []int{3, 0, 3, 4}
	expected4 := []int{3}
	actual4 := MajorityElements(input4)
	if !ArraysEqual(actual4, expected4) {
		t.Fatalf("TestMajorityElements failed. Expected: %v, got: %v", expected4, actual4)
	}
}

func TestFindMaximumXORNaive(t *testing.T) {
	input := []int{3, 10, 5, 25, 2, 8}
	expected := 28
	actual := findMaximumXORNaive(input)

	if actual != expected {
		t.Fatalf("TestFindMaximumXORNaive failed. Expected: %v, got: %v", expected, actual)
	}
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func ArraysEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

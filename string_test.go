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

func TestMyAtoi(t *testing.T) {
	input := "42"
	expected := 42
	actual := myAtoi(input)
	if expected != actual {
		t.Fatalf("MyAtoi failed. Expected: %v, got: %v", expected, actual)
	}

	input2 := "       -42"
	expected2 := -42
	actual2 := myAtoi(input2)
	if expected2 != actual2 {
		t.Fatalf("MyAtoi failed. Expected: %v, got: %v", expected2, actual2)
	}

	input3 := "4193 with words"
	expected3 := 4193
	actual3 := myAtoi(input3)
	if expected3 != actual3 {
		t.Fatalf("MyAtoi failed. Expected: %v, got: %v", expected3, actual3)
	}

	input4 := "words and 987"
	expected4 := 0
	actual4 := myAtoi(input4)
	if expected4 != actual4 {
		t.Fatalf("MyAtoi failed. Expected: %v, got: %v", expected4, actual4)
	}

	input5 := "-91283472332"
	expected5 := -2147483648
	actual5 := myAtoi(input5)
	if expected5 != actual5 {
		t.Fatalf("MyAtoi failed. Expected: %v, got: %v", expected5, actual5)
	}

	input6 := ""
	expected6 := 0
	actual6 := myAtoi(input6)
	if expected6 != actual6 {
		t.Fatalf("MyAtoi failed. Expected: %v, got: %v", expected6, actual6)
	}

	input7 := "9223372037854775808"
	expected7 := 2147483647
	actual7 := myAtoi(input7)
	if expected7 != actual7 {
		t.Fatalf("MyAtoi failed. Expected: %v, got: %v", expected7, actual7)
	}

}

package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestReverseListIterative(t *testing.T) {
	list := BuildList([]int{1, 2, 3, 4, 5})
	expected := BuildList([]int{5, 4, 3, 2, 1})
	actual := ReverseListIterative(list)
	if !(ListsEqual(expected, actual)) {
		t.Fatalf("ReverseListIterative failed, %v != %v", PrintList(expected), PrintList(actual))
	} else {
		fmt.Println("TestReverseListIterative pass")
	}
}

func TestReverseBetween(t *testing.T) {
	expected := BuildList([]int{1, 4, 3, 2, 5})
	list := BuildList([]int{1, 2, 3, 4, 5})
	actual := ReverseBetween(list, 2, 4)
	if !(ListsEqual(expected, actual)) {
		t.Fatalf("ReverseBetween failed, %v != %v", PrintList(expected), PrintList(actual))
	}

	expected2 := BuildList([]int{5})
	list2 := BuildList([]int{5})
	actual2 := ReverseBetween(list2, 1, 1)
	if !(ListsEqual(expected2, actual2)) {
		t.Fatalf("ReverseBetween failed, %v != %v", PrintList(expected2), PrintList(actual2))
	}

	expected3 := BuildList([]int{5, 3})
	list3 := BuildList([]int{3, 5})
	actual3 := ReverseBetween(list3, 1, 2)
	if !(ListsEqual(expected3, actual3)) {
		t.Fatalf("ReverseBetween failed. expected: %v, got: %v", PrintList(expected3), PrintList(actual3))
	}
	fmt.Println("TestReverseBetween pass")
}

func TestBuildList(t *testing.T) {
	expected := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	actual := BuildList([]int{1, 2, 3})
	if !(ListsEqual(expected, actual)) {
		t.Fatalf("BuildList failed, %v != %v", PrintList(expected), PrintList(actual))
	} else {
		fmt.Println("TestBuildList pass")
	}
}

func TestMakeReversedCopy(t *testing.T) {
	expected := BuildList([]int{1, 2, 3, 4})
	list := BuildList([]int{4, 3, 2, 1})
	actual := makeReversedCopy(list)

	if actual == list {
		t.Fatalf("makeReversedCopy failed, actual is not a copy of input")
	}
	if !(ListsEqual(expected, actual)) {
		t.Fatalf("makeReversedCopy failed. Expected: %v, got: %v", PrintList(expected), PrintList(actual))
	}
}
func TestMakeReversedCopy2(t *testing.T) {
	expected := BuildList([]int{1})
	list := BuildList([]int{1})
	actual := makeReversedCopy(list)

	if actual == list {
		t.Fatalf("makeReversedCopy failed, actual is not a copy of input")
	}
	if !(ListsEqual(expected, actual)) {
		t.Fatalf("makeReversedCopy failed. Expected: %v, got: %v", PrintList(expected), PrintList(actual))
	}
}

func TestAddTwoNumbers(t *testing.T) {
	list1 := BuildList([]int{7, 2, 4, 3})
	list2 := BuildList([]int{5, 6, 4})
	expected := BuildList([]int{7, 8, 0, 7})
	actual := AddTwoNumbers(list1, list2)

	if !(ListsEqual(expected, actual)) {
		t.Fatalf("AddTwoNumbers failed. Expected: %v, got: %v", PrintList(expected), PrintList(actual))
	}
}

func TestAddTwoNumbers2(t *testing.T) {
	list1 := BuildList([]int{5})
	list2 := BuildList([]int{5})
	expected := BuildList([]int{1, 0})
	actual := AddTwoNumbers(list1, list2)

	if !(ListsEqual(expected, actual)) {
		t.Fatalf("AddTwoNumbers failed. Expected: %v, got: %v", PrintList(expected), PrintList(actual))
	}
}

func TestAddTwoNumbers3(t *testing.T) {
	list1 := BuildList([]int{1})
	list2 := BuildList([]int{9, 9})
	expected := BuildList([]int{1, 0, 0})
	actual := AddTwoNumbers(list1, list2)

	if !(ListsEqual(expected, actual)) {
		t.Fatalf("AddTwoNumbers failed. Expected: %v, got: %v", PrintList(expected), PrintList(actual))
	}
}

func BuildList(a []int) *ListNode {
	if len(a) == 0 {
		return nil
	}

	head := &ListNode{
		Val: a[0],
	}
	prev := head
	var curr *ListNode

	for i := 1; i < len(a); i++ {
		curr = &ListNode{
			Val: a[i],
		}
		prev.Next = curr
		prev = curr
	}

	return head
}

func PrintList(l *ListNode) string {
	var b bytes.Buffer

	curr := l

	for curr != nil {
		nodestr := fmt.Sprintf("%v", curr.Val)
		b.WriteString(nodestr)
		b.WriteString("->")
		curr = curr.Next
	}

	b.WriteString("NULL")

	return b.String()
}

func ListsEqual(l1 *ListNode, l2 *ListNode) bool {
	if l1 == nil && l2 == nil {
		return true
	} else if l1 != nil && l2 != nil {
		return l1.Val == l2.Val && ListsEqual(l1.Next, l2.Next)
	} else {
		return false
	}
}

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
		t.Fatalf("BuidList failed, %v != %v", PrintList(expected), PrintList(actual))
	} else {
		fmt.Println("TestBuidList pass")
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

	b.WriteString("[")
	curr := l

	for curr != nil {
		nodestr := fmt.Sprintf("%v", curr.Val)
		b.WriteString(nodestr)
		if curr.Next != nil {
			b.WriteString("->")
		}
		curr = curr.Next
	}

	b.WriteString("]")

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

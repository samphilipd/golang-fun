package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
https://leetcode.com/problems/reverse-linked-list/
Reverse a singly linked list.
*/
func ReverseListIterative(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode

	for curr != nil {
		next := curr.Next
		curr.Next = prev

		prev = curr
		curr = next
	}
	return prev
}

/*
https://leetcode.com/problems/reverse-linked-list-ii/
Reverse a singly linked list from m to n
In one pass

Note: 1 ≤ m ≤ n ≤ length of list.

Example:

Input: 1->2->3->4->5->NULL, m = 2, n = 4
Output: 1->4->3->2->5->NULL

*/
func ReverseBetween(head *ListNode, m int, n int) *ListNode {
	curr := head
	var prev *ListNode

	for i := 0; i < m-1; i++ {
		prev = curr
		curr = curr.Next
	}

	if prev != nil {
		prev.Next = reverseUntil(curr, (n-m)+1)
	} else {
		head = reverseUntil(curr, n)
	}

	return head
}

func reverseUntil(head *ListNode, l int) *ListNode {
	toBeTail := head

	var prev *ListNode
	curr := head
	var next *ListNode

	for i := 0; i < l; i++ {
		next = curr.Next
		curr.Next = prev

		prev = curr
		curr = next
	}

	if toBeTail != curr {
		toBeTail.Next = curr
	}

	return prev
}

/*
You are given two non-empty linked lists representing two non-negative
integers. The most significant digit comes first and each of their nodes
contain a single digit. Add the two numbers and return it as a linked
list.

You may assume the two numbers do not contain any leading zero, except
the number 0 itself.

Follow up: What if you cannot modify the input lists? In other words,
reversing the lists is not allowed.

Example:

Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4) Output: 7 -> 8 -> 0 -> 7
*/

// func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	// reverse first
// }

func makeReversedCopy(l *ListNode) *ListNode {
	// A1 -> B1 -> C1
	// A2 <- B2
	// Build new list backwards

	var prev *ListNode
	curr := &ListNode{Val: l.Val}

	for l != nil {
		next := prev
		curr = &ListNode{Val: l.Val, Next: next}
		prev = curr
		l = l.Next
	}

	return curr
}

func length(l *ListNode) int {
	i := 0
	curr := l
	for curr != nil {
		curr = curr.Next
		i++
	}
	return i
}

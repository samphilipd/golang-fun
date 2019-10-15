/*
https://leetcode.com/problems/reverse-linked-list/
Reverse a singly linked list.

*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseListIterative(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head
	next := head

	for curr != nil {
		next = curr.Next
		curr.Next = prev

		prev = curr
		curr = next
	}
	return prev
}

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

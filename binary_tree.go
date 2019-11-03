/*
https://leetcode.com/problems/merge-two-binary-trees/

Given two binary trees and imagine that when you put one of them to cover the
other, some nodes of the two trees are overlapped while the others are not.

You need to merge them into a new binary tree. The merge rule is that if two
nodes overlap, then sum node values up as the new value of the merged node.
Otherwise, the NOT null node will be used as the node of new tree.
*/

package main

import (
	"fmt"

	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 != nil && t2 != nil {
		var newNode TreeNode
		newNode.Val = t1.Val + t2.Val
		newNode.Left = MergeTrees(t1.Left, t2.Left)
		newNode.Right = MergeTrees(t1.Right, t2.Right)
		return &newNode
	} else if t1 != nil {
		return t1
	} else if t2 != nil {
		return t2
	} else {
		return nil
	}
}

/*
226. Invert Binary Tree
Easy

Invert a binary tree.

Example:

Input:

     4
   /   \
  2     7
 / \   / \
1   3 6   9

Output:

     4
   /   \
  7     2
 / \   / \
9   6 3   1
*/
func InvertTree(root *TreeNode) *TreeNode {
	// swap the two
	// go down and swap each one
	if root == nil {
		return nil
	}
	right := InvertTree(root.Left)
	left := InvertTree(root.Right)
	root.Right = right
	root.Left = left
	return root
}

/*
101. Symmetric Tree
Easy

Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

    1
   / \
  2   2
 / \ / \
3  4 4  3



But the following [1,2,2,null,3,null,3] is not:

    1
   / \
  2   2
   \   \
   3    3



Note:
Bonus points if you could solve it both recursively and iteratively.
*/
// This is a Breadth-first search problem
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func IsSymmetricRecursive(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil {
		return false
	} else if right == nil {
		return false
	} else {
		fmt.Printf("left.val: %v\n", left.Val)
		return left.Val == right.Val &&
			isMirror(left.Left, right.Right) &&
			isMirror(left.Right, right.Left)
	}
}

func IsSymmetricIterative(root *TreeNode) bool {
	// Use a queue
	l := list.New()
	l.PushFront(root)
	l.PushFront(root)

	for l.Len() != 0 {
		n1 := dequeue(l)
		n2 := dequeue(l)
		if n1 == nil && n2 == nil {
			continue
		} else if n1 == nil || n2 == nil {
			return false
		} else if n1.Val != n2.Val {
			return false
		}
		enqueue(l, n1.Left)
		enqueue(l, n2.Right)
		enqueue(l, n1.Right)
		enqueue(l, n2.Left)
	}
	return true
}

func enqueue(l *list.List, n *TreeNode) *list.Element {
	return l.PushFront(n)
}
func dequeue(l *list.List) *TreeNode {
	el := l.Back()
	l.Remove(el)
	return el.Value.(*TreeNode)
}

/*
 Given a binary tree, you need to compute the length of the diameter of
 the tree. The diameter of a binary tree is the length of the longest
 path between any two nodes in a tree. This path may or may not pass
 through the root.

 Example: Given a binary tree

          1
         / \
        2   3
       / \
      4   5

 Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].

 Note: The length of path between two nodes is represented by the number
 of edges between them.
*/

// This solution has time complexity of O(n^2) because it must call height
// for every node
func diameterOfBinaryTree(root *TreeNode) int {
	// Use DFS
	if root == nil {
		return 0
	}

	// Passing through root
	thisDiameter := height(root.Left) + height(root.Right)

	// Not passing through root
	lDiameter := diameterOfBinaryTree(root.Left)
	rDiameter := diameterOfBinaryTree(root.Right)

	return max(thisDiameter, max(lDiameter, rDiameter))
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(height(root.Left), height(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// This implementation has time complexity of O(n) because it calculates
// height and diameter on the same pass
func diameterOfBinaryTreeOptimised(root *TreeNode) int {
	diameter, _ := diameterOfBinaryTreeOptimisedHelper(root, 0)
	return diameter
}

func diameterOfBinaryTreeOptimisedHelper(root *TreeNode, height int) (int, int) {
	// Use DFS
	if root == nil {
		return 0, 0
	}
	// Not passing through root
	lDiameter, lHeight := diameterOfBinaryTreeOptimisedHelper(root.Left, 0)
	rDiameter, rHeight := diameterOfBinaryTreeOptimisedHelper(root.Right, 0)

	// Passing through root
	thisDiameter := lHeight + rHeight
	thisHeight := max(lHeight, rHeight) + 1

	return max(thisDiameter, max(lDiameter, rDiameter)), thisHeight
}

func isBalanced(root *TreeNode) bool {
	if heightBalanced(root) == -1 {
		return false
	}
	return true
}

func heightBalanced(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lHeight := heightBalanced(root.Left)
	rHeight := heightBalanced(root.Right)

	if diff(lHeight, rHeight) <= 1 && lHeight != -1 && rHeight != -1 {
		return 1 + max(lHeight, rHeight)
	}
	return -1
}

func diff(a, b int) int {
	d := a - b
	if d < 0 {
		return -d
	}
	return d
}

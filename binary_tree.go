/*
https://leetcode.com/problems/merge-two-binary-trees/

Given two binary trees and imagine that when you put one of them to cover the
other, some nodes of the two trees are overlapped while the others are not.

You need to merge them into a new binary tree. The merge rule is that if two
nodes overlap, then sum node values up as the new value of the merged node.
Otherwise, the NOT null node will be used as the node of new tree.
*/

package main

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

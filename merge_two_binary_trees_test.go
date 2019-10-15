package main

import (
	"fmt"
	"testing"
)

func TestMergeTrees(t *testing.T) {
	tree1 := BuildTree([]int{1, 3, 2, 5})
	tree2 := BuildTree([]int{2, 1, 3, -1, 4, -1, 7})

	expected := BuildTree([]int{3, 4, 5, 5, 4, -1, 7})
	actual := MergeTrees(tree1, tree2)
	if !TreesEqual(expected, actual) {
		t.Fatalf("mergeTrees failed: trees are not equal")
	} else {
		fmt.Println("TestMergeTrees pass")
	}
}

func TestBuildTree(t *testing.T) {
	expected := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	actual := BuildTree([]int{1, 3, 2, 5})
	if !(TreesEqual(expected, actual)) {
		t.Fatalf("BuildTree failed: trees are not the same")
	} else {
		fmt.Println("TestBuildTree pass")
	}
}

func TreesEqual(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 != nil && t2 != nil {
		return t1.Val == t2.Val && TreesEqual(t1.Left, t2.Left) && TreesEqual(t1.Right, t2.Right)
	} else if t1 == nil && t2 == nil {
		return true
	} else {
		return false
	}
}

// If we observe carefully we can see that if parent node is at index i in the
// array then the left child of that node is at index (2*i + 1) and right child
// is at index (2*i + 2) in the array.
func BuildTree(vals []int) *TreeNode {
	return BuildTreeNode(vals, 0)
}

func BuildTreeNode(vals []int, i int) *TreeNode {
	if i < len(vals) {
		var root TreeNode
		val := vals[i]
		if val == -1 { // sentinel value indicating nil node {
			return nil
		}
		root.Val = vals[i]
		root.Left = BuildTreeNode(vals, (2*i)+1)
		root.Right = BuildTreeNode(vals, (2*i)+2)
		return &root

	}
	return nil
}

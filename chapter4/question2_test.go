package chapter4

import (
	"testing"
)

func Test_createMinimalBST(t *testing.T) {
	/*
			Test binary tree
			      7
			   /     \
			  3       10
			 / \     /  \
			1   4   8    13
		         \        \
			      6        14
	*/
	a := []int{1, 3, 4, 6, 7, 8, 10, 13, 14}

	want := newBinaryTreeNode(7)
	want.left = newBinaryTreeNode(3)
	want.left.left = newBinaryTreeNode(1)
	want.left.right = newBinaryTreeNode(4)
	want.left.right.right = newBinaryTreeNode(6)
	want.right = newBinaryTreeNode(10)
	want.right.left = newBinaryTreeNode(8)
	want.right.right = newBinaryTreeNode(13)
	want.right.right.right = newBinaryTreeNode(14)

	// TODO: write BST value assertion func
	got := createMinimalBST(a)
	if got == nil {
		t.Fatalf("createMinimalBST() root node is nil")
	}
	if got.value != want.value {
		t.Errorf("createMinimalBST() root node is %d want %d", got.value, want.value)
	}
	if got.left.value != want.left.value {
		t.Errorf("createMinimalBST() root.left is %d want %d", got.left.value, want.left.value)
	}
	if got.left.left.value != want.left.left.value {
		t.Errorf("createMinimalBST() root.left.left is %d want %d", got.left.left.value, want.left.left.value)
	}
	if got.left.right.value != want.left.right.value {
		t.Errorf("createMinimalBST() root.left.right is %d want %d", got.left.right.value, want.left.right.value)
	}
	if got.left.right.right.value != want.left.right.right.value {
		t.Errorf("createMinimalBST() root.left.right.right is %d want %d", got.left.right.right.value, want.left.right.right.value)
	}

	if got.right.value != want.right.value {
		t.Errorf("createMinimalBST() root.right is %d want %d", got.right.value, want.right.value)
	}
	if got.right.left.value != want.right.left.value {
		t.Errorf("createMinimalBST() root.right.left is %d want %d", got.right.left.value, want.right.left.value)
	}
	if got.right.right.value != want.right.right.value {
		t.Errorf("createMinimalBST() root.right.right is %d want %d", got.right.right.value, want.right.right.value)
	}
	if got.right.right.right.value != want.right.right.right.value {
		t.Errorf("createMinimalBST() root.right.right.right is %d want %d", got.right.right.right.value, want.right.right.right.value)
	}

}

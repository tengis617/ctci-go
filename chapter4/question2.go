package chapter4

/*
Minimal Tree:
Given a sorted (increasing order) array with unique integer elements,
write an algorithm to create a binary search tree with minimal height.
*/

type binaryTreeNode struct {
	value       int
	right, left *binaryTreeNode
}

func newBinaryTreeNode(v int) *binaryTreeNode {
	return &binaryTreeNode{value: v}
}

func createMinimalBST(a []int) *binaryTreeNode {
	return createMinimalBSTHelper(a, 0, len(a)-1)
}

func createMinimalBSTHelper(a []int, start, end int) *binaryTreeNode {
	if end < start {
		return nil
	}
	mid := (start + end) / 2

	n := newBinaryTreeNode(a[mid])
	n.left = createMinimalBSTHelper(a, start, mid-1)
	n.right = createMinimalBSTHelper(a, mid+1, end)
	return n
}

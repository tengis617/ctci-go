package chapter4

import "testing"

func Test_getList(t *testing.T) {
	root := newBinaryTreeNode(10)
	root.left = newBinaryTreeNode(1)
	root.right = newBinaryTreeNode(2)
	root.left.left = newBinaryTreeNode(4)

	lists := getList(root)

	if len(lists) != 3 {
		t.Errorf("getList() want list len to be %d got %d", 3, len(lists))
	}
	if got := lists[0].Front().Value.(int); got != 10 {
		t.Errorf("getList() level 1 head must be %d got %d", 10, got)
	}
	if got := lists[1].Front().Value.(int); got != 1 {
		t.Errorf("getList() level 1 head must be %d got %d", 1, got)
	}
	if got := lists[1].Front().Next().Value.(int); got != 2 {
		t.Errorf("getList() level 1 second element must be %d got %d", 2, got)
	}
}

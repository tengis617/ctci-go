package chapter4

import (
	"container/list"
	"fmt"
)

/*
List of Depths:
Given a binary tree, design an algorithm which creates a linked list of
all the nodes at each depth (e.g., if you have a tree with depth 0, you'll have 0 linked lists).
*/

func getList(root *binaryTreeNode) []*list.List {
	return getListHelper(root, []*list.List{}, 0)
}

func getListHelper(n *binaryTreeNode, lists []*list.List, lvl int) []*list.List {
	if n == nil {
		return lists
	}
	var ll *list.List
	if len(lists) == lvl {
		ll = list.New()
		lists = append(lists, ll)
	} else {
		ll = lists[lvl]
	}
	fmt.Println("doing", n.value)
	ll.PushBack(n.value)
	lists = getListHelper(n.left, lists, lvl+1)
	lists = getListHelper(n.right, lists, lvl+1)
	return lists
}

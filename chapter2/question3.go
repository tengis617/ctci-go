package chapter2

/*
Delete Middle Node:
Implement an algorithm to delete a node
in the middle (i.e., any node but the first and last node, not necessarily the exact middle)
of a singly linked list, given only access to that node.
*/
func deleteMiddleNode(n *node) {
	if n == nil || n.Next == nil {
		return
	}
	n.Value = n.Next.GetValue()
	n.Next = n.Next.Next
}

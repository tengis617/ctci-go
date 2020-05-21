package chapter2

import "math"

/*
Intersection:
Given two (singly) linked lists, determine if the two lists intersect.
Return the inter- secting node. Note that the intersection is defined based on reference,
not value. That is, if the kth node of the first linked list is
the exact same node (by reference) as the jth node of the second linked list,
then they are intersecting.
*/

func getIntersection(head1, head2 *node) *node {
	tail1, len1 := getTailAndLen(head1)
	tail2, len2 := getTailAndLen(head2)

	if tail1 != tail2 {
		return nil
	}

	shorter, longer := orderByLen(head1, head2, len1, len2)

	longer = getKthNode(int(math.Abs(float64(len1-len2))), longer)

	for shorter != longer {
		shorter, longer = shorter.Next, longer.Next
	}
	return longer
}

func getTailAndLen(n *node) (*node, int) {
	len := 0
	current := n
	for current.Next != nil {
		current = current.Next
		len++
	}
	return current, len
}

func orderByLen(head1, head2 *node, len1, len2 int) (shorter, longer *node) {
	if len1 > len2 {
		longer = head1
		shorter = head2
		return
	}
	longer = head2
	shorter = head1
	return
}

func getKthNode(k int, n *node) *node {
	current := n
	for i := 0; i < k; i++ {
		if current.Next == nil {
			return n
		}
		current = current.Next
	}
	return current
}

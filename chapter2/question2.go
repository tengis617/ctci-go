package chapter2

/*

Return Kth to Last:
Implement an algorithm to find the kth to last element of a singly linked list.
*/
// N => num of nodes
// Runtime: O(N)
// Space: O(1)
// returnKthToLast returns kth to last element in a singly linked list.
// Ex: k = 1 means last element in ll
// if k is less than zero, or k is out of bound, will return -1.
func returnKthToLast(ll *LinkedList, k int) int {
	if ll.head == nil || k < 1 {
		return -1
	}
	p1, p2 := ll.head, ll.head
	// p2 is runner
	for i := 0; i < k-1; i++ {
		p2 = p2.Next
		if p2 == nil {
			return -1
		}
	}

	for p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1.Value
}

package chapter2

/*
Partition:
Write code to partition a linked list around a value x,
such that all nodes less than x come before
all nodes greater than or equal to x.
If x is contained within the list,
the values of x only need to be after the elements less than x (see below).
The partition element x can appear anywhere in the "right partition";
it does not need to appear between the left and right partitions.

EXAMPLE
Input: 3 -> 5 -> 8 -> 5 -> 113 -> 2 -> 1 [partition=5]
Output: 3 -> 1 -> 2 -> 10 -> 5 -> 5 -> 8
*/

// partition partitions singly linked list by value p where all values less than p come before all
// values greater than or equal to p in unsorted order.
// N -> number of nodes in ll
// Runtime: O(N)
// Space: O(N)
func partition(ll *LinkedList, p int) *LinkedList {
	head := ll.head
	tail := ll.head

	n := ll.head
	for n != nil {
		next := n.Next
		if n.Value < p {
			n.Next = head
			head = n
		} else {
			tail.Next = n
			tail = n
		}
		n = next
	}
	tail.Next = nil
	return &LinkedList{head}
}

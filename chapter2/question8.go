package chapter2

/*
Loop Detection:
Given a circular linked list,
implement an algorithm that returns the node at the beginning of the loop.

DEFINITION
Circular linked list: A (corrupt) linked list in which a node's next pointer points
to an earlier node, so as to make a loop in the linked list.
*/

func detectLoop(ll *LinkedList) *node {
	hmap := map[*node]bool{}
	n := ll.head
	for n != nil {
		if hmap[n] {
			return n
		}
		hmap[n] = true
		n = n.Next
	}
	return nil
}

func detectLoopRunner(ll *LinkedList) *node {
	slow, fast := ll.head, ll.head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			// Collision
			break
		}
	}

	if fast == nil || fast.Next == nil {
		// not a loop
		return nil
	}

	slow = ll.head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return fast
}

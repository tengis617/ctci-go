package chapter2

/*
Remove Dups: Write code to remove duplicates from an unsorted linked list.
*/

func removeDups(ll *LinkedList) {
	exists := map[int]bool{}

	var prev *node
	current := ll.head
	for current != nil {
		if exists[current.GetValue()] {
			prev.SetNext(current.Next)
		} else {
			exists[current.GetValue()] = true
			prev = current
		}
		current = current.Next
	}
}

package chapter2

import "github.com/tengis617/ctci-go/stacks"

/*
Palindrome: Implement a function to check if a linked list is a palindrome.
*/
// isPalindrome checks if linked list is a palindrome
func isPalindrome(ll *LinkedList) bool {
	fast, slow := ll.head, ll.head
	s := stacks.NewListStack()

	for fast != nil && fast.Next != nil {
		s.Push(slow.Value)
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	for slow != nil {
		v := s.Pop()
		if v != slow.Value {
			return false
		}
		slow = slow.Next
	}
	return true
}

// BONUS: just wanted to make sure i understand this better
// findMiddleValue finds middle node value if its odd, otherwise returns -1
func findMiddleValue(ll *LinkedList) int {
	fast, slow := ll.head, ll.head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// return -1 if fast is nil
	// this means there is no "middle" node and the linked list nodes are even.
	if fast == nil {
		return -1
	}

	return slow.Value
}

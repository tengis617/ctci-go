package chapter2

/*
Sum Lists:
You have two numbers represented by a linked list, where each node contains a single digit.
The digits are stored in reverse order, such that the 1's digit is at the head of the list.
Write a function that adds the two numbers and returns the sum as a linked list.
*/

func sumList(l1, l2 *node, carry int) *node {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}

	sum := carry

	if l1 != nil {
		sum += l1.Value
	}
	if l2 != nil {
		sum += l2.Value
	}

	res := &node{Value: sum % 10}
	if l1 != nil || l2 != nil {
		next := sumList(l1.GetNext(), l2.GetNext(), getCarry(sum))
		res.SetNext(next)
	}
	return res
}

func getCarry(v int) int {
	if v >= 10 {
		return 1
	}
	return 0
}

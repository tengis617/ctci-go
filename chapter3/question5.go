package chapter3

import "github.com/tengis617/ctci-go/stacks"

/*
Sort Stack:
Write a program to sort a stack such that the smallest items are on the top.
You can use an additional temporary stack, but you may not copy the elements
into any other data structure (such as an array). The stack supports
the following operations: push, pop, peek, and isEmpty.
*/
func sortStack(s stacks.Stack) {
	tmp := stacks.NewListStack()

	for !s.IsEmpty() {
		v := s.Pop()
		for !tmp.IsEmpty() && tmp.Peek() > v {
			s.Push(tmp.Pop())
		}
		tmp.Push(v)
	}
	for !tmp.IsEmpty() {
		s.Push(tmp.Pop())
	}
}

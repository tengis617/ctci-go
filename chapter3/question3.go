package chapter3

import "github.com/tengis617/ctci-go/stacks"

/*
Stack of Plates:
Imagine a (literal) stack of plates. If the stack gets too high, it might topple.
Therefore, in real life, we would likely start a new stack when the previous stack
exceeds some threshold. Implement a data structure SetOfStacks that mimics this.
SetOfStacks should be composed of several stacks and should create a new stack once
the previous one exceeds capacity. SetOfStacks. push() and SetOfStacks. pop() should
behave identically to a single stack (that is, pop() should return the same values as
it would if there were just a single stack).

FOLLOW UP:
Implement a function popAt(int index) which performs a pop operation on a specific
sub-stack.
*/

type SetOfStacks struct {
	max int
	set []stacks.Stack
}

func newSetOfStacks(max int) *SetOfStacks {
	return &SetOfStacks{
		max: max,
		set: make([]stacks.Stack, 0),
	}
}

func (s *SetOfStacks) push(i int) {
	if s.isEmpty() || s.getTopStack().Len() >= s.max {
		s.set = append(s.set, stacks.NewListStack())
	}
	s.getTopStack().Push(i)
}

func (s *SetOfStacks) pop() int {
	if s.isEmpty() {
		return -1
	}
	v := s.getTopStack().Pop()

	if s.getTopStack().IsEmpty() {
		s.set = s.set[:len(s.set)-1]
	}
	return v
}

// NOTE: this assumes we don't need to shift from other stack when popAt() exhausts the stack.
func (s *SetOfStacks) popAt(i int) int {
	if s.isEmpty() {
		return -1
	}
	st := s.getStack(i)
	if st == nil {
		return -1
	}
	v := st.Pop()

	// remove from stack set if empty
	if st.IsEmpty() {
		s.set = append(s.set[:i], s.set[i+1:]...)
	}
	return v
}

func (s *SetOfStacks) isEmpty() bool {
	return len(s.set) == 0
}

func (s *SetOfStacks) getTopStack() stacks.Stack {
	return s.getStack(len(s.set) - 1)
}

func (s *SetOfStacks) getStack(i int) stacks.Stack {
	if i >= len(s.set) {
		return nil
	}
	return s.set[i]
}

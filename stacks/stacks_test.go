package stacks

import "testing"

func TestStack(t *testing.T) {
	stack := NewListStack()

	if !stack.IsEmpty() {
		t.Errorf("expected new stack to be empty")
	}

	stack.Push(0)
	stack.Push(1)
	stack.Push(2)

	top := stack.Peek()
	if top != 2 {
		t.Errorf("expected top to be 2, got %d", top)
	}
	stack.Pop()
	stack.Pop()

	top = stack.Peek()
	if top != 0 {
		t.Errorf("expected top to be 0, got %d", top)
	}

	if stack.IsEmpty() {
		t.Error("expected stack to not be empty")
	}

	stack.Push(11)

	if v := stack.Pop(); v != 11 {
		t.Errorf("expected top to be 11, got %d", v)
	}
	stack.Pop()
	if !stack.IsEmpty() {
		t.Error("expected stack to be empty")
	}

	if stack.Peek() != -1 {
		t.Error("expected nil")
	}

	if stack.Pop() != -1 {
		t.Error("expected -1")
	}
}

package stacks

import "testing"

func TestStack(t *testing.T) {
	stack := NewListStack()

	if !stack.IsEmpty() {
		t.Errorf("expected new stack to be empty")
	}

	stack.Push(&Node{0})
	stack.Push(&Node{1})
	stack.Push(&Node{2})

	top := stack.Peek()
	if top.Value != 2 {
		t.Errorf("expected top to be 2, got %d", top.Value)
	}
	stack.Pop()
	stack.Pop()

	top = stack.Peek()
	if top.Value != 0 {
		t.Errorf("expected top to be 0, got %d", top)
	}

	if stack.IsEmpty() {
		t.Error("expected stack to not be empty")
	}

	stack.Push(&Node{11})

	if v := stack.Pop(); v.Value != 11 {
		t.Errorf("expected top to be 11, got %d", v.Value)
	}
	stack.Pop()
	if !stack.IsEmpty() {
		t.Error("expected stack to be empty")
	}

	if stack.Peek() != nil {
		t.Error("expected nil")
	}

	if stack.Pop() != nil {
		t.Error("expected nil")
	}
}

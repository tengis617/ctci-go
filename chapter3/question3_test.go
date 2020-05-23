package chapter3

import (
	"testing"
)

func Test_SetOfStacks(t *testing.T) {
	ss := newSetOfStacks(3)

	// first stack starts from 0
	for i := 0; i < 3; i++ {
		ss.push(i)
	}
	// second stack starts from 10
	ss.push(10)
	if len(ss.set) != 2 {
		t.Errorf("should have %d stacks, got %d", 2, len(ss.set))
	}

	if v := ss.pop(); v != 10 {
		t.Errorf("should return %d, got %d", 10, v)
	}

	if len(ss.set) != 1 {
		t.Errorf("should have %d stacks, got %d", 1, len(ss.set))
	}
	if got := ss.set[len(ss.set)-1].Len(); got != ss.max {
		t.Errorf("want %d, got %d", ss.max, got)
	}

	// second stack starts from 10
	for i := 0; i < 3; i++ {
		ss.push(10 + i)
	}

	// third stack starts from 20
	for i := 0; i < 3; i++ {
		ss.push(20 + i)
	}
	if len(ss.set) != 3 {
		t.Errorf("should have %d stacks, got %d", 3, len(ss.set))
	}

	// pop second stack until empty
	for i := 0; i < 3; i++ {
		if got := ss.popAt(2); got != 22-i {
			t.Errorf("should return %d got %d", 22-i, got)
		}
	}
	if len(ss.set) != 2 {
		t.Errorf("should have %d stacks, got %d", 2, len(ss.set))
	}

	// try popping at out of range index
	if got := ss.popAt(10); got != -1 {
		t.Errorf("should return %d got %d", -1, got)
	}

	// pop remaining
	for i := 0; i < 6; i++ {
		ss.pop()
	}
	if !ss.isEmpty() {
		t.Errorf("should be empty after popping all")
	}
	if got := ss.pop(); got != -1 {
		t.Errorf("should return -1 when empty, got %d", got)
	}
	if got := len(ss.set); got != 0 {
		t.Errorf("should not have any sets, got %d", got)
	}

	// try popAt on empty stack set
	if got := ss.popAt(1); got != -1 {
		t.Errorf("should return %d got %d", -1, got)
	}
}

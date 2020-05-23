package chapter3

import (
	"testing"
)

func Test_newMyQueue(t *testing.T) {
	// keep track of step for testing
	q := newMyQueue()
	if !q.IsEmpty() {
		t.Error("queue must be empty")
	}
	if got := q.Remove(); got != -1 {
		t.Error("Remove() must return -1 when empty")
	}
	q.Add(1)
	if q.IsEmpty() {
		t.Error("q must not be empty after adding")
	}
	if got := q.Peek(); got != 1 {
		t.Errorf("Peek() want %d got %d", 1, got)
	}

	for i := 2; i <= 10; i++ {
		q.Add(i)
	}
	for i := 1; i <= 10; i++ {
		if got := q.Remove(); got != i {
			t.Errorf("Remove() want %d got %d", i, got)
		}
	}
	if !q.IsEmpty() {
		t.Error("should be empty after Remove() all")
	}

	if got := q.Remove(); got != -1 {
		t.Errorf("Remove() should return -1 when empty got %d", got)
	}

	if got := q.Peek(); got != -1 {
		t.Errorf("Peek() should return -1 when empty got %d", got)
	}
}

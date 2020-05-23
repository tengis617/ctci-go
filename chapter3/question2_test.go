package chapter3

import (
	"testing"
)

func Test_Min(t *testing.T) {
	s := newMinStack()
	s.Push(2)
	if min := s.Min(); min != 2 {
		t.Errorf("expected min to be %d got %d", 2, min)
	}
	s.Push(4)
	if min := s.Min(); min != 2 {
		t.Errorf("expected min to be %d got %d", 2, min)
	}
	s.Push(1)
	if min := s.Min(); min != 1 {
		t.Errorf("expected min to be %d got %d", 1, min)
	}
	if v := s.Pop(); v != 1 {
		t.Errorf("expected pop to return %d got %d", 1, v)
	}
	if min := s.Min(); min != 2 {
		t.Errorf("expected min to be %d got %d", 2, min)
	}
	s.Push(6)
	if min := s.Min(); min != 2 {
		t.Errorf("expected min to be %d got %d", 2, min)
	}
	s.Push(1)
	if min := s.Min(); min != 1 {
		t.Errorf("expected min to be %d got %d", 1, min)
	}
	if v := s.Pop(); v != 1 {
		t.Errorf("expected pop to return %d got %d", 1, v)
	}
	if v := s.Pop(); v != 6 {
		t.Errorf("expected pop to return %d got %d", 6, v)
	}
	if v := s.Pop(); v != 4 {
		t.Errorf("expected pop to return %d got %d", 4, v)
	}
	if v := s.Pop(); v != 2 {
		t.Errorf("expected pop to return %d got %d", 2, v)
	}
	if v := s.Pop(); v != -1 {
		t.Errorf("expected empty stack pop to return %d got %d", -1, v)
	}
	if min := s.Min(); min != -1 {
		t.Errorf("expected min to be %d got %d", -1, min)
	}
}

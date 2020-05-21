package chapter2

import (
	"reflect"
	"testing"
)

func Test_detectLoop(t *testing.T) {
	e := &node{Value: 1}
	d := &node{Value: 1, Next: e}
	c := &node{Value: 1, Next: d}
	b := &node{Value: 1, Next: c}
	a := &node{Value: 1, Next: b}
	// circular
	e.Next = c

	e1 := &node{Value: 1}
	d1 := &node{Value: 1, Next: e1}
	c1 := &node{Value: 1, Next: d1}
	b1 := &node{Value: 1, Next: c1}
	a1 := &node{Value: 1, Next: b1}
	// circular
	e1.Next = a1

	tests := []struct {
		name string
		ll   *LinkedList
		want *node
	}{
		{"a -> b -> c -> d -> e -> c", &LinkedList{head: a}, c},
		{"a -> b -> c -> d -> e -> a", &LinkedList{head: a1}, a1},
		{"no loop should return nil", NewLinkedList(1, 2, 3, 4), nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectLoop(tt.ll); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectLoop() = %v, want %v", got, tt.want)
			}
			if got := detectLoopRunner(tt.ll); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectLoopRunner() = %v, want %v", got, tt.want)
			}
		})
	}
}

package chapter2

import (
	"reflect"
	"testing"
)

func Test_getIntersection(t *testing.T) {
	/*
		a -> b -> c -> d -> g -> h
			 e -> f -> d
	*/

	h := &node{Value: 1}
	g := &node{Value: 1, Next: h}
	d := &node{Value: 1, Next: g}
	c := &node{Value: 1, Next: d}
	b := &node{Value: 1, Next: c}
	a := &node{Value: 1, Next: b}

	f := &node{Value: 1, Next: d}
	e := &node{Value: 1, Next: f}

	tests := []struct {
		name  string
		head1 *node
		head2 *node
		want  *node
	}{
		{"a -> b -> c -> d -> g -> h and  e -> f -> d ", a, e, d},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersection(tt.head1, tt.head2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

package chapter3

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/tengis617/ctci-go/stacks"
)

func Test_sortStack(t *testing.T) {
	s := stacks.NewListStack()
	s.Push(2)
	s.Push(6)
	s.Push(8)
	s.Push(4)
	s.Push(1)

	sorted := stacks.NewListStack()
	sorted.Push(8)
	sorted.Push(6)
	sorted.Push(4)
	sorted.Push(2)
	sorted.Push(1)

	tests := []struct {
		s    stacks.Stack
		want stacks.Stack
	}{
		{s, sorted},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			if sortStack(tt.s); !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("sortStack() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}

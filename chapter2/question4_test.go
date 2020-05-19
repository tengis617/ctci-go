package chapter2

import (
	"fmt"
	"testing"
)

func Test_partition(t *testing.T) {
	tests := []struct {
		ll        *LinkedList
		partition int
		want      *LinkedList
	}{
		// not the best way to test
		{NewLinkedList(3, 5, 8, 5, 10, 2, 1), 5, NewLinkedList(1, 2, 3, 5, 8, 5, 10)},
		{NewLinkedList(3, 3, 3, 1, 1, 1), 2, NewLinkedList(1, 1, 1, 3, 3, 3)},
		{NewLinkedList(1, 3, 1, 3, 1, 3), 2, NewLinkedList(1, 1, 1, 3, 3, 3)},
		{NewLinkedList(1, 2, 3), 2, NewLinkedList(1, 2, 3)},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("ll:%s  partition: %d", tt.ll.print(), tt.partition), func(t *testing.T) {
			got := partition(tt.ll, tt.partition)
			if got.print() != tt.want.print() {
				t.Errorf("partition() = %v, want %v", got.print(), tt.want.print())
			}
		})
	}
}

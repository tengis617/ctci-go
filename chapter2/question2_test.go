package chapter2

import (
	"fmt"
	"testing"
)

func Test_returnKthToLast(t *testing.T) {
	tests := []struct {
		ll   *LinkedList
		k    int
		want int
	}{
		{NewLinkedList(0, 1, 2, 3, 4, 5), 3, 3},
		{NewLinkedList(0, 1, 2, 3, 4, 5), 6, 0},
		{NewLinkedList(0, 1, 2, 3, 4), 2, 3},
		{NewLinkedList(0, 1, 2), 2, 1},
		{NewLinkedList(0, 1), 2, 0},
		{NewLinkedList(0), 1, 0},
		// edge cases
		{NewLinkedList(0), 2, -1},
		{NewLinkedList(0, 1, 2, 3, 4), 0, -1},
		{NewLinkedList(0, 1, 2, 3, 4), 10, -1},
		{NewLinkedList(0, 1, 2, 3, 4, 5), 7, -1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("ll: %s k: %d", tt.ll.print(), tt.k), func(t *testing.T) {
			if got := returnKthToLast(tt.ll, tt.k); got != tt.want {
				t.Errorf("returnKthToLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

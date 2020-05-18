package chapter2

import (
	"fmt"
	"testing"
)

func Test_deleteMiddleNode(t *testing.T) {
	tests := []struct {
		ll        *LinkedList
		nodeIndex int
		want      *LinkedList
	}{
		{NewLinkedList(1, 2, 3, 4, 5), 2, NewLinkedList(1, 2, 4, 5)},
		{NewLinkedList(1, 2, 3, 4, 5), 3, NewLinkedList(1, 2, 3, 5)},
		{NewLinkedList(1, 2, 3), 1, NewLinkedList(1, 3)},

		// should not change if last element
		{NewLinkedList(1, 2, 3, 4, 5), 4, NewLinkedList(1, 2, 3, 4, 5)},
		// should not change if out of range
		{NewLinkedList(1, 2, 3, 4, 5), 5, NewLinkedList(1, 2, 3, 4, 5)},
	}
	for _, tt := range tests {
		nodeToDelete := tt.ll.head
		for i := 0; i < tt.nodeIndex; i++ {
			nodeToDelete = nodeToDelete.Next
		}
		t.Run(fmt.Sprintf("ll:%s  nodeIndex: %d", tt.ll.print(), tt.nodeIndex), func(t *testing.T) {
			deleteMiddleNode(nodeToDelete)
			if tt.ll.print() != tt.want.print() {
				t.Errorf("deleteMiddleNode() = %v, want %v", tt.ll.print(), tt.want.print())
			}
		})
	}
}

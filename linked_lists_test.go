package ctcigo

import (
	"fmt"
	"testing"
)

func TestRemoveDupsHashMap(t *testing.T) {
	tests := []struct {
		l    *LinkedList
		want *LinkedList
	}{
		{NewLinkedList(1, 2, 3, 3, 4), NewLinkedList(1, 2, 3, 4)},
		{NewLinkedList(1, 1, 1, 1, 1), NewLinkedList(1)},
		{NewLinkedList(1, 0, 1, 0, 1), NewLinkedList(1, 0)},
		{NewLinkedList(1, 1), NewLinkedList(1)},
		{NewLinkedList(1, 2, 2, 2, 2), NewLinkedList(1, 2)},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt.l.removeDupsHashMap()

			// neat hack to do value comparison
			if tt.l.print() != tt.want.print() {
				t.Errorf("got %v want %v", tt.l.print(), tt.want.print())
			}
		})
	}
}

func TestRemoveDupsRunner(t *testing.T) {
	tests := []struct {
		l    *LinkedList
		want *LinkedList
	}{
		{NewLinkedList(1, 2, 3, 3, 4), NewLinkedList(1, 2, 3, 4)},
		{NewLinkedList(1, 1, 1, 1, 1), NewLinkedList(1)},
		{NewLinkedList(1, 0, 1, 0, 1), NewLinkedList(1, 0)},
		{NewLinkedList(1, 1), NewLinkedList(1)},
		{NewLinkedList(1, 2, 2, 2, 2), NewLinkedList(1, 2)},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt.l.removeDupsRunner()

			// neat hack to do value comparison
			if tt.l.print() != tt.want.print() {
				t.Errorf("got %v want %v", tt.l.print(), tt.want.print())
			}
		})
	}
}

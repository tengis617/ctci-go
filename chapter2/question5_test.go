package chapter2

import (
	"fmt"
	"testing"
)

func Test_sumList(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		l1   *LinkedList
		l2   *LinkedList
		want *LinkedList
	}{
		{
			NewLinkedList(7, 1, 6),
			NewLinkedList(5, 9, 2),
			NewLinkedList(2, 1, 9),
		},
		{
			NewLinkedList(7, 7, 7),
			NewLinkedList(2, 2, 2),
			NewLinkedList(9, 9, 9),
		},
		{
			NewLinkedList(6, 6, 6),
			NewLinkedList(4, 4, 4),
			NewLinkedList(0, 1, 1, 1),
		},
		{
			NewLinkedList(6),
			NewLinkedList(4),
			NewLinkedList(0, 1),
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("l1: %s l2: %s want: %s", tt.l1.print(), tt.l2.print(), tt.want.print()), func(t *testing.T) {
			gotNode := sumList(tt.l1.head, tt.l2.head, 0)
			gotList := &LinkedList{head: gotNode}
			if gotList.print() != tt.want.print() {
				t.Errorf("sumList() = %v, want %v", gotList.print(), tt.want.print())
			}
		})
	}
}

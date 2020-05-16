package chapter2

import (
	"reflect"
	"testing"
)

func Test_removeDups(t *testing.T) {
	tests := []struct {
		ll   *LinkedList
		want *LinkedList
	}{
		{NewLinkedList(1, 2, 3), NewLinkedList(1, 2, 3)},
		{NewLinkedList(1, 2, 3, 3), NewLinkedList(1, 2, 3)},
		{NewLinkedList(1, 1, 1, 1), NewLinkedList(1)},
	}
	for _, tt := range tests {
		t.Run(tt.ll.print(), func(t *testing.T) {
			removeDups(tt.ll)
			if !reflect.DeepEqual(tt.ll, tt.want) {
				t.Errorf("removeDups() = %s want %s", tt.ll.print(), tt.want.print())
			}
		})
	}
}

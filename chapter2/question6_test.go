package chapter2

import (
	"reflect"
	"testing"
)

func Test_findMiddle(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		ll   *LinkedList
		want int
	}{
		{NewLinkedList(1), 1},
		{NewLinkedList(1, 2), -1},
		{NewLinkedList(1, 2, 3), 2},
		{NewLinkedList(1, 2, 3, 4), -1},
		{NewLinkedList(1, 2, 3, 4, 5), 3},
		{NewLinkedList(1, 2, 3, 4, 5, 6), -1},
		{NewLinkedList(1, 2, 3, 4, 5, 6, 7), 4},
	}
	for _, tt := range tests {
		t.Run(tt.ll.print(), func(t *testing.T) {
			if got := findMiddleValue(tt.ll); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMiddle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		ll   *LinkedList
		want bool
	}{
		{NewLinkedList(1, 2, 3, 2, 1), true},
		{NewLinkedList(1, 2, 3, 4, 4, 3, 2, 1), true},
		{NewLinkedList(1, 1, 1, 1, 1), true},
		{NewLinkedList(1, 1, 1, 1), true},
		{NewLinkedList(1), true},
		{NewLinkedList(1, 1), true},
		{NewLinkedList(1, 2), false},
		{NewLinkedList(1, 2, 3, 4), false},
		{NewLinkedList(1, 1, 1, 1, 1, 0), false},
	}
	for _, tt := range tests {
		t.Run(tt.ll.print(), func(t *testing.T) {
			if got := isPalindrome(tt.ll); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

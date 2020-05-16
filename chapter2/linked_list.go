package chapter2

import (
	"fmt"
	"strings"
)

// node represents a singly linked list node
type node struct {
	Value int
	Next  *node
}

func (n *node) SetNext(next *node) {
	n.Next = next
}

func (n *node) GetValue() int {
	return n.Value
}

type LinkedList struct {
	head *node
}

func (ll *LinkedList) append(v int) {
	newNode := &node{Value: v}
	if ll.head == nil {
		ll.head = newNode
		return
	}

	current := ll.head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}

func (ll *LinkedList) print() string {
	if ll.head == nil {
		return ""
	}
	var b strings.Builder
	n := ll.head
	for n != nil {
		b.WriteString(fmt.Sprintf("%d", n.Value))
		if n.Next != nil {
			b.WriteString(",")
		}
		n = n.Next
	}
	return b.String()
}

// LinkedList creates a new singly linked list and returns the head node
func NewLinkedList(values ...int) *LinkedList {
	ll := &LinkedList{}
	for _, v := range values {
		ll.append(v)
	}
	return ll
}

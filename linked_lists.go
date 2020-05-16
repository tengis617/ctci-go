package ctcigo

import (
	"fmt"
	"strings"
)

type Node struct {
	Next *Node
	Data int
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Append(d int) {
	end := &Node{Data: d}

	if l.Head == nil {
		l.Head = end
		return
	}

	n := l.Head
	for n.Next != nil {
		n = n.Next
	}

	n.Next = end
}

func (l *LinkedList) print() string {
	var b strings.Builder
	n := l.Head
	for n != nil {
		b.WriteString(fmt.Sprintf("%d,", n.Data))
		n = n.Next
	}
	return b.String()
}

func NewLinkedList(data ...int) *LinkedList {
	l := &LinkedList{}
	for _, d := range data {
		l.Append(d)
	}
	return l
}

// removeDupsHashMap removes duplicates in a linked list using hash maps
// runs in O(n)
// takes 0(n) space
func (l *LinkedList) removeDupsHashMap() {
	exists := make(map[int]bool)
	var previous *Node

	n := l.Head
	for n != nil {
		if exists[n.Data] {
			previous.Next = n.Next
		} else {
			exists[n.Data] = true
			previous = n
		}
		n = n.Next
	}
}

// removeDupsHashMap removes duplicates in a linked list using two pointers
// runs in O(n^2)
// takes 0(1) space
func (l *LinkedList) removeDupsRunner() {
	current := l.Head
	for current != nil {
		runner := current
		for runner.Next != nil {
			if runner.Next.Data == current.Data {
				runner.Next = runner.Next.Next
			} else {
				runner = runner.Next
			}
		}
		current = current.Next
	}
}

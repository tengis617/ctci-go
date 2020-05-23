package chapter3

import (
	"github.com/tengis617/ctci-go/queue"
	"github.com/tengis617/ctci-go/stacks"
)

/*
Queue via Stacks:
Implement a MyQueue class which implements a queue using two stacks.
*/

type myQueue struct {
	new stacks.Stack
	old stacks.Stack
}

// newMyQueue creates a queue using 2 stacks.
func newMyQueue() queue.Queue {
	return &myQueue{
		new: stacks.NewListStack(),
		old: stacks.NewListStack(),
	}
}

func (q *myQueue) Add(i int) {
	q.new.Push(i)
}

func (q *myQueue) Remove() int {
	if q.IsEmpty() {
		return -1
	}
	q.shift()
	return q.old.Pop()
}

func (q *myQueue) Peek() int {
	if q.IsEmpty() {
		return -1
	}
	q.shift()
	return q.old.Peek()
}

func (q *myQueue) IsEmpty() bool {
	return q.new.IsEmpty() && q.old.IsEmpty()
}

func (q *myQueue) shift() {
	if q.old.IsEmpty() {
		for !q.new.IsEmpty() {
			q.old.Push(q.new.Pop())
		}
	}
}

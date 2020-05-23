package stacks

type Stack interface {
	Pop() int
	Push(n int)
	Peek() int
	IsEmpty() bool
	Len() int
}

type ListStack struct {
	list  []int
	count int
}

// NewListStack creates a new stack using a list.
func NewListStack() Stack {
	return &ListStack{
		list: make([]int, 0),
	}
}

// Push pushes a new element onto the stack.
func (s *ListStack) Push(n int) {
	s.list = append(s.list[:s.count], n)
	s.count++
}

// Pop pops the topmost element on the stack.
func (s *ListStack) Pop() int {
	n := s.Peek()
	s.count--
	return n
}

// IsEmpty returns true if the stack is empty.
func (s *ListStack) IsEmpty() bool {
	return s.count == 0
}

func (s *ListStack) Peek() int {
	if s.IsEmpty() {
		return -1
	}
	return s.list[s.count-1]
}

func (s *ListStack) Len() int {
	return s.count
}

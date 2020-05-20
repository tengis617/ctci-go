package stacks

type Stack interface {
	Pop() int
	Push(n int)
	Peek() int
	IsEmpty() bool
}

type ListStack struct {
	list  []int
	count int
}

func NewListStack() Stack {
	return &ListStack{
		list: make([]int, 0),
	}
}

func (s *ListStack) Push(n int) {
	s.list = append(s.list[:s.count], n)
	s.count++
}

func (s *ListStack) Pop() int {
	n := s.Peek()
	s.count--
	return n
}

func (s *ListStack) IsEmpty() bool {
	return s.count == 0
}

func (s *ListStack) Peek() int {
	if s.IsEmpty() {
		return -1
	}
	return s.list[s.count-1]
}

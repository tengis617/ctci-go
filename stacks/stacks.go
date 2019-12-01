package stacks

type Stack interface {
	Pop() *Node
	Push(n *Node)
	Peek() *Node
	IsEmpty() bool
}

type Node struct {
	Value int
}

type ListStack struct {
	nodes []*Node
	count int
}

func NewListStack() Stack {
	return &ListStack{
		nodes: make([]*Node, 0),
	}
}

func (s *ListStack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *ListStack) Pop() *Node {
	n := s.Peek()
	s.count--
	return n
}

func (s *ListStack) IsEmpty() bool {
	return s.count == 0
}

func (s *ListStack) Peek() *Node {
	if s.IsEmpty() {
		return nil
	}
	return s.nodes[s.count-1]
}

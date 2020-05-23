package chapter3

/*
Stack Min:
How would you design a stack which,
in addition to push and pop, has a function min
which returns the minimum element? Push, pop and min
should all operate in 0(1) time.
*/

type minStack struct {
	list []int
	min  []int
}

func newMinStack() *minStack {
	return &minStack{
		list: make([]int, 0),
		min:  make([]int, 0),
	}
}

func (s *minStack) Push(i int) {
	s.list = append(s.list, i)
	if len(s.min) == 0 {
		s.min = append(s.min, i)
		return
	}
	if i < s.Min() {
		s.min = append(s.min, i)
	}
}

func (s *minStack) Pop() int {
	if len(s.list) == 0 {
		return -1
	}
	v := s.list[len(s.list)-1]
	if s.list[len(s.list)-1] == s.Min() {
		s.min = s.min[:len(s.min)-1]
	}
	s.list = s.list[:len(s.list)-1]
	return v
}

func (s *minStack) Min() int {
	if len(s.min) == 0 {
		return -1
	}
	return s.min[len(s.min)-1]
}

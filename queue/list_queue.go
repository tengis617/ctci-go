package queue

type Queue interface {
	Add(int)
	Remove() int
	Peek() int
	IsEmpty() bool
}

type queue struct {
	list  []int
	count int
}

func NewListQueue() Queue {
	return &queue{
		list: make([]int, 0),
	}
}

func (q *queue) Add(i int) {
	q.list = append([]int{i}, q.list...)
	q.count++
}

func (q *queue) Remove() int {
	if q.IsEmpty() {
		return -1
	}
	v := q.list[len(q.list)-1]
	q.list = q.list[:len(q.list)-1]
	q.count--
	return v
}

func (q *queue) Peek() int {
	if q.IsEmpty() {
		return -1
	}
	return q.list[0]
}

func (q *queue) IsEmpty() bool {
	return q.count == 0
}

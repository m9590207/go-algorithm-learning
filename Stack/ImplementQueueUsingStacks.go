package Stack

type MyStack struct {
	data []int
}

func (s *MyStack) Push(v int) {
	s.data = append(s.data, v)
}

func (s *MyStack) Pop() int {
	l := len(s.data) - 1
	res := s.data[l]
	s.data = s.data[:l]
	return res
}

func (s *MyStack) Peek() int {
	return s.data[len(s.data)-1]
}

func (s *MyStack) Empty() bool {
	return len(s.data) == 0
}

type MyQueue struct {
	in  MyStack
	out MyStack
}

func Constructor() MyQueue {
	q := MyQueue{}
	q.in.data = []int{}
	q.out.data = []int{}
	return q
}

func (q *MyQueue) Push(v int) {
	q.in.Push(v)
}

func (q *MyQueue) Pop() int {
	for !q.in.Empty() {
		q.out.Push(q.in.Pop())
	}
	front := q.out.Pop()
	for !q.out.Empty() {
		q.in.Push(q.out.Pop())
	}
	return front
}

func (q *MyQueue) Peek() int {
	for !q.in.Empty() {
		q.out.Push(q.in.Pop())
	}
	front := q.out.Peek()
	for !q.out.Empty() {
		q.in.Push(q.out.Pop())
	}
	return front
}

func (q *MyQueue) Empty() bool {
	return q.in.Empty()
}

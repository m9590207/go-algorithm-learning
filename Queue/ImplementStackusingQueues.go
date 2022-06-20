package Queue

type queue struct {
	data []int
}

func initQueue() queue {
	return queue{data: make([]int, 0)}
}

func (q *queue) Enqueue(v int) {
	q.data = append(q.data, v)
}

func (q *queue) Dequeue() int {
	v := q.data[0]
	q.data = q.data[1:]
	return v
}

func (q *queue) Front() int {
	return q.data[0]
}

func (q *queue) IsEmpty() bool {
	return len(q.data) == 0
}

type MyStack struct {
	q queue
}

func Constructor() MyStack {
	s := MyStack{}
	s.q = initQueue()
	return s
}

func (this *MyStack) Push(x int) {
	this.q.Enqueue(x)
	for i := 0; i < len(this.q.data)-1; i++ {
		this.q.Enqueue(this.q.Dequeue())
	}
}

func (this *MyStack) Pop() int {
	return this.q.Dequeue()
}

func (this *MyStack) Top() int {
	return this.q.Front()
}

func (this *MyStack) Empty() bool {
	return this.q.IsEmpty()
}

package Queue

type MyCircularQueue struct {
	frontIndex int
	size       int
	arr        []int
}

func InitMyCircularQueue(k int) MyCircularQueue {
	queue := MyCircularQueue{
		frontIndex: 0,
		size:       0,
		arr:        make([]int, k),
	}
	return queue
}

func (q *MyCircularQueue) EnQueue(value int) bool {
	n := len(q.arr)
	if q.size == n {
		return false
	}

	nextIndex := (q.frontIndex + q.size) % n
	q.arr[nextIndex] = value
	q.size++
	return true
}

func (q *MyCircularQueue) DeQueue() bool {
	if q.size == 0 {
		return false
	}

	q.frontIndex = (q.frontIndex + 1) % len(q.arr)
	q.size--
	return true
}

func (q *MyCircularQueue) Front() int {
	if q.size == 0 {
		return -1
	}
	return q.arr[q.frontIndex]
}

func (q *MyCircularQueue) Rear() int {
	if q.size == 0 {
		return -1
	}
	return q.arr[(q.frontIndex+q.size-1)%len(q.arr)]
}

func (q *MyCircularQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *MyCircularQueue) IsFull() bool {
	return q.size == len(q.arr)
}

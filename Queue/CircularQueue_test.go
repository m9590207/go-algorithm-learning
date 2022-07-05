package Queue

import "testing"

func TestCircularQueueEnqueue(t *testing.T) {
	q := InitMyCircularQueue(5)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	if q.Front() != 1 {
		t.Errorf("Enquque 錯誤")
	}
}

func TestCircularQueueDequeue(t *testing.T) {
	q := InitMyCircularQueue(5)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.DeQueue()
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	q.DeQueue()
	q.EnQueue(7)
	if q.Front() != 3 {
		t.Errorf("Dequque 錯誤")
	}
}

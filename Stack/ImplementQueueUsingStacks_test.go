package Stack

import "testing"

func TestMyQueuePush(t *testing.T) {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	if q.Peek() != 1 {
		t.Errorf("Push 錯誤")
	}
}

func TestMyQueuePop(t *testing.T) {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	if q.Pop() != 1 {
		t.Errorf("Pop 錯誤")
	}
}

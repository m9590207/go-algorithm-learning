package Queue

import (
	"testing"
)

func TestEnqueue(t *testing.T) {
	CustomQueue := &customQueue{}
	CustomQueue.Enqueue("A")
	CustomQueue.Enqueue("B")
	CustomQueue.Enqueue("C")

	if CustomQueue.queue[0] != "A" || CustomQueue.queue[2] != "C" {
		t.Errorf("Enqueue 錯誤: 資料結果與預期不符")
	}
}
func TestDequeue(t *testing.T) {
	customQueue := &customQueue{}
	customQueue.Enqueue("A")
	customQueue.Enqueue("B")
	customQueue.Enqueue("C")
	err := customQueue.Dequeue()
	if err != nil {
		t.Errorf("Dequeue 錯誤: %s", err.Error())
	}

	if customQueue.queue[0] != "B" {
		t.Errorf("Dequeue 錯誤: 未依FIFO順序執行")
	}
}
func TestFront(t *testing.T) {
	customQueue := &customQueue{}
	customQueue.Enqueue("A")
	customQueue.Enqueue("B")
	customQueue.Enqueue("C")
	name, err := customQueue.Front()

	if err != nil {
		t.Errorf("Front 錯誤: %s", err.Error())
	}

	if name != "A" {
		t.Errorf("Front 錯誤: 非第一個傳入元素")
	}

}
func TestSize(t *testing.T) {
	customQueue := &customQueue{}
	customQueue.Enqueue("A")
	customQueue.Enqueue("B")
	customQueue.Enqueue("C")

	if customQueue.Size() != 3 {
		t.Errorf("Size 錯誤")
	}
}
func TestEmpty(t *testing.T) {
	customQueue := &customQueue{}
	customQueue.Enqueue("A")

	if customQueue.Empty() {
		t.Errorf("Empty 錯誤")
	}
}

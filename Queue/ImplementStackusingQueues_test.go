package Queue

import "testing"

func TestPush(t *testing.T) {
	s := Constructor()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if s.Top() != 3 {
		t.Errorf("Push 錯誤top != 3")
	}
}

func TestPop(t *testing.T) {
	s := Constructor()
	s.Push(1)
	s.Push(2)
	if s.Pop() != 2 {
		t.Errorf("Pop 錯誤 != 2")
	}
}

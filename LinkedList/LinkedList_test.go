package LinkedList

import "testing"

func TestAddFront(t *testing.T) {
	customList := &linkedList{}
	customList.AddFront(1)
	customList.AddFront(2)
	customList.AddFront(3)

	if customList.length != 3 {
		t.Errorf("AddFront 錯誤: length 不正確")
	}

	if customList.head.data != 3 {
		t.Errorf("AddFront 錯誤: head 不正確")
	}
}

func TestRemoveFront(t *testing.T) {
	customList := &linkedList{}
	customList.AddFront(1)
	customList.AddFront(2)

	if err := customList.RemoveFront(); err != nil {
		t.Errorf("RemoveFront 錯誤: %s", err.Error())
	}
	if customList.length != 1 {
		t.Errorf("RemoveFront 錯誤: length 不正確")
	}
	if customList.head.data != 1 {
		t.Errorf("RemoveFront 錯誤: head 不正確")
	}
}

func TestAddBack(t *testing.T) {
	customList := &linkedList{}
	customList.AddBack(1)
	customList.AddBack(2)
	customList.AddBack(3)

	if customList.length != 3 {
		t.Errorf("AddBack 錯誤: length 不正確")
	}

	if customList.head.data != 1 {
		t.Errorf("AddBack 錯誤: head 不正確")
	}
}

func TestRemoveBack(t *testing.T) {
	customList := &linkedList{}
	customList.AddFront(1)
	customList.AddFront(2)

	if err := customList.RemoveBack(); err != nil {
		t.Errorf("RemoveBack 錯誤: %s", err.Error())
	}
	if customList.length != 1 {
		t.Errorf("RemoveBack 錯誤: length 不正確")
	}
	if customList.head.data != 2 {
		t.Errorf("RemoveBack 錯誤: head 不正確")
	}
}

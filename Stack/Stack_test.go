package Stack

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	customStack := &customStack{}
	customStack.Push("A")
	customStack.Push("B")
	customStack.Push("C")

	if customStack.Empty() {
		fmt.Errorf("Push 錯誤stack為空")
	}
	if len(customStack.stack) != 3 {
		fmt.Errorf("Push 錯誤stack長度不符")
	}
}

func TestPop(t *testing.T) {
	customStack := &customStack{}
	customStack.Push("A")
	customStack.Push("B")
	customStack.Push("C")

	if err := customStack.Pop(); err != nil {
		fmt.Errorf("Pop 錯誤: %s", err.Error())
	}
	len := len(customStack.stack)
	if len != 2 {
		fmt.Errorf("Pop 錯誤stack長度不符")
	}

	if customStack.stack[len-1] != "B" {
		fmt.Errorf("Pop 錯誤最上層的元素不正確")
	}
}

func TestPeek(t *testing.T) {
	customStack := &customStack{}
	customStack.Push("A")
	customStack.Push("B")
	customStack.Push("C")

	top, err := customStack.Peek()
	if err != nil {
		fmt.Errorf("Peek 錯誤: %s", err.Error())
	}
	if top != "C" {
		fmt.Errorf("Peek 錯誤最上層的元素不正確")
	}
}

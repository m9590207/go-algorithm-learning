package Stack

import (
	"fmt"
	"sync"
)

type customStack struct {
	stack []string
	lock  sync.RWMutex
}

func (c *customStack) Push(value string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.stack = append(c.stack, value)
}

func (c *customStack) Pop() error {
	c.lock.Lock()
	defer c.lock.Unlock()
	len := len(c.stack)
	if len > 0 {
		c.stack = c.stack[:len-1]
		return nil
	}
	return fmt.Errorf("Pop 錯誤 stack為空")
}

func (c *customStack) Peek() (string, error) {
	if len(c.stack) > 0 {
		return c.stack[0], nil
	}
	return "", fmt.Errorf("Peek 錯誤 stack為空")
}

func (c *customStack) Size() int {
	return len(c.stack)
}

func (c *customStack) Empty() bool {
	return len(c.stack) == 0
}

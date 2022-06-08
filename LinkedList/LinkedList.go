package LinkedList

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) AddFront(value int) {
	n := &node{data: value}
	if l.head == nil {
		l.head = n
	} else {
		n.next = l.head
		l.head = n
	}
	l.length++
}

func (l *linkedList) RemoveFront() error {
	if l.head == nil {
		return fmt.Errorf("List為空")
	}
	l.head = l.head.next
	l.length--
	return nil
}

func (l *linkedList) AddBack(value int) {
	n := &node{data: value}
	if l.head == nil {
		l.head = n
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = n
	}
	l.length++
}

func (l *linkedList) RemoveBack() error {
	if l.head == nil {
		return fmt.Errorf("List為空")
	}
	current := l.head
	var prev *node
	for current.next != nil {
		prev = current
		current = current.next
	}
	if prev != nil {
		prev.next = nil
	} else {
		l.head = nil
	}
	l.length--
	return nil
}

func (l *linkedList) Fornt() (int, error) {
	if l.head == nil {
		return 0, fmt.Errorf("List為空")
	}
	return l.head.data, nil
}

func (l *linkedList) Size() int {
	return l.length
}

func (l *linkedList) Traverse() error {
	if l.head == nil {
		return fmt.Errorf("List為空")
	}
	current := l.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
	return nil
}

package main

import (
	"fmt"

	"github.com/m9590207/go-algorithm-learning/BinarySearch"
)

func main() {
	l1n1 := &ListNode{Val: 3}
	l1n2 := &ListNode{Val: 4, Next: l1n1}
	l1n3 := &ListNode{Val: 2, Next: l1n2}
	l1n4 := &ListNode{Val: 7, Next: l1n3}

	l2n1 := &ListNode{Val: 4}
	l2n2 := &ListNode{Val: 6, Next: l2n1}
	l2n3 := &ListNode{Val: 5, Next: l2n2}

	result := addTwoNumbers(l1n4, l2n3)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}

	fmt.Println("BinarySearch", BinarySearch.Search([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := initStack(l1)
	s2 := initStack(l2)
	preHead := &ListNode{}
	carry := 0
	for {
		if len(s1) == 0 && len(s2) == 0 && carry == 0 {
			break
		}
		v1 := popValue(&s1)
		v2 := popValue(&s2)
		sum := v1 + v2 + carry
		node := &ListNode{Val: sum % 10}
		node.Next = preHead.Next
		preHead.Next = node
		carry = sum / 10
	}
	return preHead.Next
}

func initStack(head *ListNode) []int {
	var stack []int
	for {
		if head == nil {
			break
		}
		stack = append(stack, head.Val)
		head = head.Next
	}
	return stack
}

func popValue(stack *[]int) int {
	length := len(*stack)
	var lastValue int
	if length > 0 {
		lastIndex := length - 1
		lastValue = (*stack)[lastIndex]
		*stack = (*stack)[:lastIndex]
	}
	return lastValue
}

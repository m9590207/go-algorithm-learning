package BinarySearchTree

import "fmt"

type node struct {
	data  int
	left  *node
	right *node
}

type bst struct {
	root *node
}

func (b *bst) Reset() {
	b.root = nil
}

func (b *bst) insertRec(n *node, data int) *node {
	if b.root == nil {
		b.root = &node{data: data}
		return b.root
	}
	if n == nil {
		return &node{data: data}
	}
	if data <= n.data {
		n.left = b.insertRec(n.left, data)
	}
	if data > n.data {
		n.right = b.insertRec(n.right, data)
	}
	return n

}

func (b *bst) Insert(data int) {
	b.insertRec(b.root, data)
}

func (b *bst) findRec(n *node, data int) *node {
	if n == nil {
		return nil
	}
	if n.data == data {
		return n
	}
	if n.data <= data {
		b.findRec(n.left, data)
	}
	return b.findRec(n.right, data)
}

func (b *bst) Find(data int) bool {
	n := b.findRec(b.root, data)
	if n == nil {
		return false
	} else {
		return true
	}
}

func (b *bst) inOrderRec(n *node) {
	if n == nil {
		return
	}
	b.inOrderRec(n.left)
	fmt.Println(n.data)
	b.inOrderRec(n.right)
}

func (b *bst) InOrder() {
	b.inOrderRec(b.root)
}

package main

import "fmt"

type Node[T any] struct {
	data  T
	left  *Node[T]
	right *Node[T]
}

func NewNode[T any](data T) *Node[T] {
	node := Node[T]{
		data:  data,
		left:  nil,
		right: nil,
	}

	return &node
}

func (n Node[T]) String() string {
	return fmt.Sprintf("Data : %v left : %v right :%v\n", n.data, n.left, n.right)
}

type Tree[T any] struct {
	root *Node[T]
}

func NewTree[T any](root *Node[T]) *Tree[T] {
	tree := &Tree[T]{
		root: root,
	}

	return tree
}

func updateStringPointer(storage *string, value string) {

	tempo := *storage + value

	*storage = tempo
}

func (t Tree[T]) inorder(node *Node[T], storage *string) {

	if node != nil {
		t.inorder(node.left, storage)
		updateStringPointer(storage, fmt.Sprintf("%v -> ", node.data))
		t.inorder(node.right, storage)
	}

}

func (t Tree[T]) preorder(node *Node[T], storage *string) {

	if node != nil {
		updateStringPointer(storage, fmt.Sprintf("%v -> ", node.data))
		t.preorder(node.left, storage)
		t.preorder(node.right, storage)
	}

}

func (t Tree[T]) postorder(node *Node[T], storage *string) {

	if node != nil {

		t.postorder(node.left, storage)
		t.postorder(node.right, storage)
		updateStringPointer(storage, fmt.Sprintf("%v -> ", node.data))
	}

}

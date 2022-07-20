package binarytree

import (
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	data  T
	left  *Node[T]
	right *Node[T]
}

func NewNode[T constraints.Ordered](data T) *Node[T] {
	node := Node[T]{
		data:  data,
		left:  nil,
		right: nil,
	}

	return &node
}

func (n Node[T]) String() string {
	return fmt.Sprintf("\nData : %v left : %v right :%v\n", n.data, n.left, n.right)
}

func (n Node[T]) hasOneChild() bool {
	return (n.left == nil && n.right != nil) || (n.left != nil && n.right == nil)
}

func (n Node[T]) hasTwoChild() bool {
	return n.left != nil && n.right != nil
}

func (n Node[T]) isALeaf() bool {
	return n.left == nil && n.right == nil
}

//The Tree struct is not needed, but was really helpful in my mind :)
type Tree[T constraints.Ordered] struct {
	root *Node[T]
}

func NewTree[T constraints.Ordered](root *Node[T]) *Tree[T] {
	tree := &Tree[T]{
		root: root,
	}

	return tree
}

//NR -> None recursive
func (t Tree[T]) SearchNR(looking T) bool {

	node := t.root
	flag := false

	if node == nil {
		return flag
	}

	for !flag && node != nil {

		if looking == node.data {
			return true
		} else if looking < node.data {
			node = node.left
		} else {
			node = node.right
		}
	}

	return flag

}

//R -> Recursive
func (t Tree[T]) SearchR(looking T) *Node[T] {

	//Because i have created Tree struct
	var search func(node *Node[T], looking T) *Node[T]

	search = func(node *Node[T], looking T) *Node[T] {
		if node == nil {
			return nil
		}

		if node.data == looking {
			return node
		}

		if looking < node.data {
			return search(node.left, looking)
		}

		return search(node.right, looking)
	}

	return search(t.root, looking)

}

func (t Tree[T]) InsertNR(insert T) {

	node := t.root

	flag := false

	for !flag {

		if insert < node.data {
			if node.left == nil {
				node.left = NewNode(insert)
				flag = true
			}
			node = node.left
		} else {
			if node.right == nil {
				node.right = NewNode(insert)
				flag = true
			}
			node = node.right
		}
	}

}

func (t Tree[T]) InsertR(value T) {

	var insert func(node *Node[T], value T)

	insert = func(node *Node[T], value T) {
		if value < node.data {
			if node.left == nil {
				node.left = NewNode(value)
				return
			}
			insert(node.left, value)
		} else {
			if node.right == nil {
				node.right = NewNode(value)
				return
			}
			insert(node.right, value)
		}

	}

	insert(t.root, value)

}

func (t Tree[T]) CountNodeR() int {

	var res int = 0

	var traversal func(node *Node[T])

	traversal = func(node *Node[T]) {
		if node != nil {
			res++
			traversal(node.left)
			traversal(node.right)
		}
	}

	traversal(t.root)

	return res

}

//TODO : upate code
func (t Tree[T]) Delete(value T) {

	if t.SearchR(value) != nil {

		var parent, node *Node[T] = nil, t.root
		flag := false

		if t.IsEmpty() {
			return
		}

		for !flag {

			if value < node.data {
				parent = node
				node = node.left
			} else if value > node.data {
				parent = node
				node = node.right
			} else {

				if node.isALeaf() {
					//Case leaf node

					if parent.left.data == node.data {
						parent.left = nil
					} else {
						parent.right = nil
					}

				} else if node.hasOneChild() {
					//Case one node has one child
					var childOfNode *Node[T] = nil

					if node.left != nil {
						childOfNode = node.left
					} else {
						childOfNode = node.right
					}

					//Side of previous node
					if parent.left.data == node.data {
						parent.left = childOfNode
					} else {
						parent.right = childOfNode
					}
				} else {
					//Case left two child

					i := t.Inorder(node)

					t.Delete(i.data)

					node.data = i.data

				}

				flag = true

			}
		}

	}

}

func (t Tree[T]) GetParent(node *Node[T], value T) *Node[T] {

	if node == nil {
		return nil
	}

	var previous *Node[T] = nil

	for node != nil {
		if value < node.data {
			previous = node
			node = node.left
		} else if value > node.data {
			previous = node
			node = node.right
		} else {
			node = nil
		}
	}

	return previous
}

func (t Tree[T]) PeekRoot() *Node[T] {
	return t.root
}

func (t Tree[T]) IsEmpty() bool {
	return t.root == nil
}

func (t Tree[T]) Inorder(node *Node[T]) *Node[T] {

	if t.IsEmpty() {
		return nil
	}

	tree := node.right
	for tree.left != nil {
		tree = tree.left
	}
	return tree

}

func updateStringPointer(storage *string, value string) {

	tempo := *storage + value

	*storage = tempo
}

func (t Tree[T]) Display(node *Node[T], level int) {

	if node != nil {
		t.Display(node.left, level+1)
		str := fmt.Sprintf("%s -> %v", strings.Repeat(" ", 4*level), node.data)
		fmt.Println(str)
		t.Display(node.right, level+1)
	}

}

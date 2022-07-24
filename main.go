package main

import (
	"github.com/Michelprogram/data-structures/binarytree"
	"github.com/Michelprogram/data-structures/stack"
	"github.com/Michelprogram/data-structures/linkedlist"
)

func main() {

	tree := binarytree.NewTree(binarytree.NewNode(1))
	tree.Display(tree.PeekRoot(), 0)

	stack.NewStack[int](1).IsEmpty()

	linkedlist.NewLinkedList[]()
}

package main

import "github.com/Michelprogram/data-structures/binarytree"

func main() {

	tree := binarytree.NewTree(binarytree.NewNode(1))
	tree.Display(tree.PeekRoot(), 0)
}

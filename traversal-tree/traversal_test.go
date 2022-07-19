package main

import (
	"testing"
)

func BasicTree() *Tree[int] {

	tree := NewTree(NewNode(1))

	tree.root.left = NewNode(2)
	tree.root.right = NewNode(3)

	tree.root.left.left = NewNode(4)
	tree.root.left.right = NewNode(5)

	tree.root.left.right.left = NewNode(21)

	return tree
}

func TestUpdatePointer(t *testing.T) {

	var storage string = "Hello"

	updateStringPointer(&storage, " World")

	if storage != "Hello World" {
		t.Errorf("Update pointer was incorrect, got: %s, want: %v.", storage, "Hello World")
	}
}

func TestNode(t *testing.T) {

	node := NewNode(1)

	node.left = NewNode(2)

	node.left.left = NewNode(3)

	if node.right != nil {
		t.Errorf("Node right of root node was incorrect, got: %v, want: %v.", node.right, nil)

	}

	if node.left.left.data != 3 {
		t.Errorf("Third level Node was incorrect, got: %v, want: %v.", node.left.left.data, 3)
	}

}

func TestTree(t *testing.T) {

	tree := NewTree(NewNode(1))

	if tree.root.data != 1 {
		t.Errorf("Tree root Node was incorrect, got: %v, want: %v.", tree.root.data, 1)

	}

}

func TestInorder(t *testing.T) {

	var res string = ""

	tree := BasicTree()

	tree.inorder(tree.root, &res)

	expected := "4 -> 2 -> 21 -> 5 -> 1 -> 3 -> "

	if res != expected {
		t.Errorf("Inorder was incorrect, got: %s, want: %s.", res, expected)
	}

}

func TestPreorder(t *testing.T) {
	var res string = ""

	tree := BasicTree()

	tree.preorder(tree.root, &res)

	expected := "1 -> 2 -> 4 -> 5 -> 21 -> 3 -> "

	if res != expected {
		t.Errorf("Inorder was incorrect, got: %s, want: %s.", res, expected)
	}
}

func TestPostorder(t *testing.T) {
	var res string = ""

	tree := BasicTree()

	tree.postorder(tree.root, &res)

	expected := "4 -> 21 -> 5 -> 2 -> 3 -> 1 -> "

	if res != expected {
		t.Errorf("Inorder was incorrect, got: %s, want: %s.", res, expected)
	}
}

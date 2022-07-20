package binarytree

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

func PerfectBinaryTree() *Tree[int] {

	tree := NewTree(NewNode(8))

	tree.root.left = NewNode(4)

	tree.root.left.left = NewNode(2)

	tree.root.left.left.left = NewNode(1)
	tree.root.left.left.right = NewNode(3)

	tree.root.left.right = NewNode(6)

	tree.root.left.right.right = NewNode(7)
	tree.root.left.right.left = NewNode(5)

	tree.root.right = NewNode(12)

	tree.root.right.left = NewNode(10)

	tree.root.right.left.left = NewNode(9)
	tree.root.right.left.right = NewNode(11)

	tree.root.right.right = NewNode(14)

	tree.root.right.right.right = NewNode(15)
	tree.root.right.right.left = NewNode(13)

	return tree

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

func TestNodeIsLeaf(t *testing.T) {

	tree := PerfectBinaryTree()

	get := tree.SearchR(1)

	if !get.isALeaf() {
		t.Errorf("Is leaf was incorrect, got: %t, want: %t.", get.isALeaf(), true)
	}

	get = tree.SearchR(2)

	if get.isALeaf() {
		t.Errorf("Is leaf was incorrect, got: %t, want: %t.", get.isALeaf(), false)
	}

}

func TestNodeHasOneChild(t *testing.T) {

	tree := PerfectBinaryTree()

	tree.InsertR(21)

	get := tree.SearchR(15)

	if !get.hasOneChild() {
		t.Errorf("Has one child was incorrect, got: %t, want: %t.", get.hasOneChild(), true)
	}

	get = tree.SearchR(2)

	if get.hasOneChild() {
		t.Errorf("has one child was incorrect, got: %t, want: %t.", get.hasOneChild(), false)
	}

}

func TestNodeHasTwoChild(t *testing.T) {

	tree := PerfectBinaryTree()

	get := tree.SearchR(14)

	if !get.hasTwoChild() {
		t.Errorf("Has two child was incorrect, got: %t, want: %t.", get.hasOneChild(), true)
	}

	get = tree.SearchR(2)

	if get.hasOneChild() {
		t.Errorf("has two child was incorrect, got: %t, want: %t.", get.hasOneChild(), false)
	}

}

func TestTree(t *testing.T) {

	tree := NewTree(NewNode(1))

	if tree.root.data != 1 {
		t.Errorf("Tree root Node was incorrect, got: %v, want: %v.", tree.root.data, 1)
	}

}

func TestSearchNR(t *testing.T) {

	tree := PerfectBinaryTree()

	if !tree.SearchNR(1) {
		t.Errorf("Search NR was incorrect, got: %t, want: %t.", tree.SearchNR(1), true)
	}

	if tree.SearchNR(21) {
		t.Errorf("Search NR was incorrect, got: %t, want: %t.", tree.SearchNR(21), false)
	}
}

func TestPeekRoot(t *testing.T) {

	tree := PerfectBinaryTree()

	if tree.PeekRoot().data != 8 {
		t.Errorf("Peek root was incorrect, got: %v, want: %v.", tree.PeekRoot().data, 8)
	}

}

func TestSearchR(t *testing.T) {

	tree := PerfectBinaryTree()

	if tree.SearchR(1) == nil {
		t.Errorf("Search R was incorrect, got: %v, want: %s.", tree.SearchR(1), "not nil")
	}

	if tree.SearchR(21) != nil {
		t.Errorf("Search R was incorrect, got: %v, want: %v.", tree.SearchR(21), nil)
	}
}

func TestInsertNR(t *testing.T) {

	tree := PerfectBinaryTree()

	tree.InsertNR(21)

	if tree.SearchR(21) == nil {
		t.Errorf("InsertNR was incorrect, got: %v, want: %s.", tree.SearchR(21), "not nil")
	}

}

func TestInsertR(t *testing.T) {

	tree := PerfectBinaryTree()

	tree.InsertR(21)

	if tree.SearchR(21) == nil {
		t.Errorf("Insert R was incorrect, got: %v, want: %v.", tree.SearchR(21), "not nil")
	}

}

func TestCountNode(t *testing.T) {

	tree := PerfectBinaryTree()

	lenBase := tree.CountNodeR()

	if lenBase != 15 {
		t.Errorf("Count was incorrect, got: %d, want: %d.", lenBase, 15)
	}

	tree.InsertR(21)

	if tree.CountNodeR() != lenBase+1 {
		t.Errorf("Count was incorrect, got: %d, want: %d.", tree.CountNodeR(), lenBase+1)

	}

}

func TestParentNode(t *testing.T) {

	tree := PerfectBinaryTree()

	get := tree.GetParent(tree.root, 3).data

	if get != 2 {
		t.Errorf("Get parent was incorrect, got: %v, want: %d.", get, 2)
	}

}

func TestRemoveNodeLeaf(t *testing.T) {

	tree := PerfectBinaryTree()

	tree.Delete(1)

	if tree.SearchR(1) != nil {
		t.Errorf("Remove leaf node was incorrect, got: %v, want: %v.", tree.SearchR(1), nil)
	}

}

func TestRemoveNodeOneChild(t *testing.T) {

	tree := PerfectBinaryTree()

	tree.InsertR(21)

	tree.Delete(15)

	if tree.SearchR(14).right.data != 21 {
		t.Errorf("Remove node with one child was incorrect, got: %v, want: %v.", tree.SearchR(14).right.data, 21)
	}

}

func TestRemoveNodeOneChild2(t *testing.T) {

	tree := PerfectBinaryTree()

	tree.InsertR(21)
	tree.InsertR(19)

	tree.InsertR(25)
	tree.InsertR(23)

	tree.Delete(15)

	if tree.SearchR(14).right.right.data != 25 {
		t.Errorf("Remove node with one child was incorrect, got: %v, want: %v.", tree.SearchR(14).right.right.data, 25)
	}

}

func TestRemoveNodeTwoChild(t *testing.T) {

	tree := PerfectBinaryTree()

	tree.Delete(10)

	if tree.SearchR(10) != nil {
		t.Errorf("Remove node with two child was incorrect, got: %v, want: %v.", tree.SearchR(10), nil)
	}

}

func TestRemoveNodeTwoChild2(t *testing.T) {

	tree := PerfectBinaryTree()

	tree.Delete(12)

	if tree.SearchR(14).left != nil {
		t.Errorf("Remove node with two child was incorrect, got: %v, want: %v.", tree.SearchR(12), nil)
	}

}

func TestDispaly(t *testing.T) {

	tree := PerfectBinaryTree()

	tree.Display(tree.root, 0)

}

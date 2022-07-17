package main

import (
	"testing"
)

func initLinkedList() LinkedList[int] {
	node1, node2, node3 := NewNode(1, nil), NewNode(2, nil), NewNode(3, nil)

	linkedList := NewLinkedList[int]()

	linkedList.head = node1

	node1.next = node2

	node2.next = node3

	return *linkedList
}

func unsortedLinkedList() LinkedList[int] {
	node1, node2, node3 := NewNode(5, nil), NewNode(1, nil), NewNode(3, nil)

	linkedList := NewLinkedList[int]()

	linkedList.head = node1

	node1.next = node2

	node2.next = node3

	return *linkedList
}

func TestNewNode(t *testing.T) {

	node1, node2, tail := NewNode(1, nil), NewNode(2, nil), NewNode(3, nil)

	node1.next = node2

	node2.next = tail

	head := NewNode(1, node1)

	if head.next.data != 1 {
		t.Errorf("Node was incorrect, got: %d, want: %d.", head.next.data, 1)

	}

}

func TestNewLinkedList(t *testing.T) {

	linkedList := initLinkedList()

	if linkedList.head.data != 1 {
		t.Errorf("Head was incorrect, got: %d, want: %d.", linkedList.head.data, 1)
	}

}

func TestTraversal(t *testing.T) {

	linkedList := initLinkedList()

	way := linkedList.Traversal()
	expected := "1 --> 2 --> 3\n"

	if way != expected {
		t.Errorf("Traversal was incorrect, got: %s, want: %s.", way, expected)

	}

}

func TestPickFirst(t *testing.T) {

	linkedList := initLinkedList()

	if linkedList.PeekFirst().data != 2 {
		t.Errorf("Insertion was incorrect, got: %d, want: %d.", linkedList.PeekFirst().data, 2)
	}

}

func TestPickLast(t *testing.T) {

	linkedList := initLinkedList()

	if linkedList.PeekLast().data != 3 {
		t.Errorf("Insertion was incorrect, got: %d, want: %d.", linkedList.PeekLast().data, 3)
	}

}

func TestInsertBegin(t *testing.T) {

	linkedList := initLinkedList()

	linkedList.InsertionBeginning(21)

	if linkedList.PeekHead().data != 21 {
		t.Errorf("Insertion was incorrect, got: %d, want: %d.", linkedList.PeekFirst().data, 21)
	}

}

func TestInsertEnd(t *testing.T) {

	linkedList := initLinkedList()

	linkedList.InsertionEnding(21)

	if linkedList.PeekLast().data != 21 {
		t.Errorf("Insertion was incorrect, got: %d, want: %d.", linkedList.PeekLast().data, 21)
	}

}

func TestSearch(t *testing.T) {

	linkedList := initLinkedList()

	res := linkedList.Search(1)

	if res == nil {
		t.Errorf("Search was incorrect, got: %v, want: %v.", res, linkedList.PeekFirst())
	}

	res = linkedList.Search(21)

	if res != nil {
		t.Errorf("Search was incorrect, got: %v, want: %v.", res, nil)
	}

}

func TestDeleteBegin(t *testing.T) {

	linkedList := initLinkedList()

	linkedList.DeleteBeginning()

	if linkedList.PeekHead().data != 2 {
		t.Errorf("Delete beign was incorrect, got: %d, want: %d.", linkedList.PeekFirst().data, 2)

	}

}

func TestDeleteEnding(t *testing.T) {

	linkedList := initLinkedList()

	linkedList.DeleteEnding()

	if linkedList.PeekLast().data != 2 {
		t.Errorf("Delete end was incorrect, got: %d, want: %d.", linkedList.PeekLast().data, 2)

	}

}

func TestDelete(t *testing.T) {

	linkedList := initLinkedList()

	linkedList.Delete(2)

	if linkedList.PeekLast().data != 3 {
		t.Errorf("Delete was incorrect, got: %d, want: %d.", linkedList.PeekLast().data, 3)

	}

}

func TestSort(t *testing.T) {

	linkedList := unsortedLinkedList()

	linkedList.Sort()

	expected := "1 --> 3 --> 5\n"

	if linkedList.Traversal() != expected {
		t.Errorf("Sort was incorrect, got: %s, want: %s.", linkedList.Traversal(), expected)

	}

}

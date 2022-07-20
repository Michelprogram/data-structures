package linkedlist

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Node[T comparable] struct {
	data T
	next *Node[T]
}

func NewNode[T constraints.Ordered](data T, next *Node[T]) *Node[T] {
	node := Node[T]{
		data: data,
		next: next,
	}

	return &node
}

func (n Node[T]) String() string {
	return fmt.Sprintf("Data : %v Next : %v\n", n.data, n.next.data)
}

// Display list
func (n Node[T]) Display(occurence int) string {

	fmt.Printf("Data : %v Index : %d\n", n.data, occurence)

	if n.next == nil {
		return ""
	}

	occurence++

	return n.next.Display(occurence)
}

type LinkedList[T constraints.Ordered] struct {
	head *Node[T]
}

func NewLinkedList[T constraints.Ordered]() *LinkedList[T] {

	linkedList := LinkedList[T]{}

	return &linkedList
}

//Access each element of the linked list
func (l *LinkedList[T]) Traversal() string {

	var res string
	var node *Node[T] = l.head

	for node.next != nil {

		res += fmt.Sprintf("%v --> ", node.data)

		node = node.next
	}

	res += fmt.Sprintf("%v\n", node.data)

	return res
}

//Adds an element at the beginning of the list.
func (l *LinkedList[T]) InsertionBeginning(element T) {

	node := NewNode(element, l.head)

	l.head = node

}

//Adds an element at the ending of the list.
func (l *LinkedList[T]) InsertionEnding(element T) {

	node := NewNode(element, nil)

	l.PeekLast().next = node

}

//Deletes an element at the beginning of the list.

func (l *LinkedList[T]) DeleteBeginning() {
	l.head = l.head.next
}

//Deletes an element using the given key.
func (l *LinkedList[T]) DeleteEnding() {

	var node *Node[T] = l.head

	for node.next.next != nil {

		node = node.next
	}

	node.next = nil

}

//Deletes an element using the given key.
func (l *LinkedList[T]) Delete(element T) *Node[T] {

	item := l.Search(element)

	if item != nil {

		var node *Node[T] = l.head

		for node.next != item {

			node = node.next
		}

		node.next = node.next.next

		return node

	}

	return item

}

//Searches an element using the given key.
func (l *LinkedList[T]) Search(element T) *Node[T] {

	var node *Node[T] = l.head

	for node.next != nil {

		if node.data == element {
			return node
		}

		node = node.next
	}

	return nil
}

//Return the first element after head
func (l *LinkedList[T]) PeekFirst() *Node[T] {
	return l.head.next
}

//Return head
func (l *LinkedList[T]) PeekHead() *Node[T] {
	return l.head
}

//Return the last element
func (l *LinkedList[T]) PeekLast() *Node[T] {

	var node *Node[T] = l.head

	for node.next != nil {

		node = node.next
	}

	return node
}

//Sort LinkedList

func (l *LinkedList[T]) Sort() {

	var node *Node[T] = l.head
	var index *Node[T] = nil

	if l.head == nil {
		return
	}

	for node.next != nil {

		index = node.next

		if node.data > index.data {
			node.data, index.data = index.data, node.data
		}

		node = node.next
	}

}

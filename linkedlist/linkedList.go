package linkedlist

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Node[T comparable] struct {
	data T
	Next *Node[T]
}

func NewNode[T constraints.Ordered](data T, Next *Node[T]) *Node[T] {
	node := Node[T]{
		data: data,
		Next: Next,
	}

	return &node
}

func (n Node[T]) String() string {
	return fmt.Sprintf("Data : %v Next : %v\n", n.data, n.Next.data)
}

// Display list
func (n Node[T]) Display(occurence int) string {

	fmt.Printf("Data : %v Index : %d\n", n.data, occurence)

	if n.Next == nil {
		return ""
	}

	occurence++

	return n.Next.Display(occurence)
}

type LinkedList[T constraints.Ordered] struct {
	Head *Node[T]
}

func NewLinkedList[T constraints.Ordered]() *LinkedList[T] {

	linkedList := LinkedList[T]{}

	return &linkedList
}

//Access each element of the linked list
func (l *LinkedList[T]) Traversal() string {

	var res string
	var node *Node[T] = l.Head

	if node == nil {
		return ""
	}

	for node.Next != nil {

		res += fmt.Sprintf("%v --> ", node.data)

		node = node.Next
	}

	res += fmt.Sprintf("%v\n", node.data)

	return res
}

//Adds an element at the beginning of the list.
func (l *LinkedList[T]) InsertionBeginning(element T) {

	node := NewNode(element, l.Head)

	l.Head = node

}

//Adds an element at the ending of the list.
func (l *LinkedList[T]) InsertionEnding(element T) {

	node := NewNode(element, nil)

	l.PeekLast().Next = node

}

//Deletes an element at the beginning of the list.

func (l *LinkedList[T]) DeleteBeginning() {
	l.Head = l.Head.Next
}

//Deletes an element using the given key.
func (l *LinkedList[T]) DeleteEnding() {

	var node *Node[T] = l.Head

	for node.Next.Next != nil {

		node = node.Next
	}

	node.Next = nil

}

//Deletes an element using the given key.
func (l *LinkedList[T]) Delete(element T) *Node[T] {

	item := l.Search(element)

	if item != nil {

		var node *Node[T] = l.Head

		for node.Next != item {

			node = node.Next
		}

		node.Next = node.Next.Next

		return node

	}

	return item

}

//Searches an element using the given key.
func (l *LinkedList[T]) Search(element T) *Node[T] {

	var node *Node[T] = l.Head

	for node.Next != nil {

		if node.data == element {
			return node
		}

		node = node.Next
	}

	return nil
}

//Return the first element after Head
func (l *LinkedList[T]) PeekFirst() *Node[T] {
	return l.Head.Next
}

//Return Head
func (l *LinkedList[T]) PeekHead() *Node[T] {
	return l.Head
}

//Return the last element
func (l *LinkedList[T]) PeekLast() *Node[T] {

	var node *Node[T] = l.Head

	for node.Next != nil {

		node = node.Next
	}

	return node
}

//Sort LinkedList

func (l *LinkedList[T]) Sort() {

	var node *Node[T] = l.Head
	var index *Node[T] = nil

	if l.Head == nil {
		return
	}

	for node.Next != nil {

		index = node.Next

		if node.data > index.data {
			node.data, index.data = index.data, node.data
		}

		node = node.Next
	}

}

func (l LinkedList[T]) Size() int {

	if l.Head == nil {
		return 0
	}

	i := 0

	node := l.Head

	for node.Next != nil {
		i++
		node = node.Next
	}
	return i

}

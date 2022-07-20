package stack

import "fmt"

type Stack[T any] struct {
	arr  []T
	top  int
	size int
}

func NewStack[T any](size int) *Stack[T] {

	stack := Stack[T]{
		arr:  make([]T, size),
		top:  0,
		size: size,
	}

	return &stack

}

//Add an element to the top of a stack

func (s *Stack[T]) Push(element T) (T, error) {

	if s.IsFull() {

		return nilItem[T](), fmt.Errorf("The stack is full !")
	}

	s.arr[s.top] = element

	s.top++

	return element, nil
}

//Remove an element from the top of a stack
func (s *Stack[T]) Pop() (T, error) {

	if s.IsEmpty() {

		return nilItem[T](), fmt.Errorf("The stack is empty !")
	}

	item := s.arr[s.top-1]

	s.top--

	return item, nil
}

//Check if the stack is empty

func (s Stack[T]) IsEmpty() bool {
	return s.top == 0
}

//Check if the stack is full
func (s Stack[T]) IsFull() bool {
	return s.top == s.size
}

//Get the value of the top element without removing it
func (s Stack[T]) Peek() T {
	return s.arr[s.top-1]
}

//Show what we have inside Stack
func (s Stack[T]) String() string {
	var res string

	for index, element := range s.arr {
		res += fmt.Sprintf("Index : %d Value : %v\n", index, element)
	}

	return res
}

func nilItem[T any]() T {
	var item T
	return item
}

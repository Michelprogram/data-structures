package queue

import "fmt"

type Queue[T any] struct {
	arr   []T
	size  int
	front int
	rear  int
}

func NewQueue[T any](size int) *Queue[T] {

	queue := Queue[T]{
		arr:   make([]T, size),
		front: -1,
		rear:  -1,
		size:  size,
	}

	return &queue

}

func (q *Queue[T]) String() string {
	var res string

	for index, val := range q.arr {

		res += fmt.Sprintf("Index : %d Value : %v\n", index, val)
	}

	return res
}

// Add an element to the end of the queue
func (q *Queue[T]) Enqueue(element T) (T, error) {

	if q.IsFull() {
		return nilValue[T](), fmt.Errorf("The queue is full !")
	}

	if q.front == -1 {
		q.front = 0
	}
	q.rear++

	q.arr[q.rear] = element

	return element, nil

}

//Remove an element from the front of the queue
func (q *Queue[T]) Dequeue() (T, error) {

	if q.IsEmpty() {
		return nilValue[T](), fmt.Errorf("The queue is empty !")
	}

	item := q.arr[q.front]

	for i := 0; i < q.rear; i++ {
		q.arr[i] = q.arr[i+1]
	}
	q.rear--

	if q.rear == -1 {
		q.front = -1
	}

	return item, nil

}

// Check if the queue is empty
func (q Queue[T]) IsEmpty() bool {
	return q.front == -1 && q.rear == -1
}

// Check if the queue is full
func (q Queue[T]) IsFull() bool {
	return q.rear == q.size-1
}

//	Get the value of the front of the queue without removing it
func (q Queue[T]) Peek() T {
	return q.arr[q.front]
}

func nilValue[T any]() T {
	var item T
	return item
}

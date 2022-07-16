package main

import (
	"testing"
)

func TestNewQueue(t *testing.T) {

	queueInt := NewQueue[int](3)

	if len(queueInt.arr) != 3 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", queueInt.size, 3)
	}

	queueStr := NewQueue[string](5)

	if len(queueStr.arr) != 5 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", queueStr.size, 5)
	}
}

func TestEnqueueDequeue(t *testing.T) {

	queueInt := NewQueue[int](5)

	queueInt.Enqueue(1)
	queueInt.Enqueue(2)
	queueInt.Enqueue(3)
	queueInt.Enqueue(4)
	queueInt.Enqueue(5)

	_, err := queueInt.Enqueue(6)

	if err == nil {
		t.Errorf("Enqueue was incorrect, expected error got nil")

	}

	if queueInt.Peek() != 1 {
		t.Errorf("Peek was incorrect, got: %d, want: %d.", queueInt.Peek(), 1)
	}

	queueInt.Dequeue()
	queueInt.Dequeue()

	if queueInt.Peek() != 3 {
		t.Errorf("Peek was incorrect, got: %d, want: %d.", queueInt.Peek(), 3)

	}
}

func TestIsFull(t *testing.T) {
	queueInt := NewQueue[int](2)

	queueInt.Enqueue(1)
	queueInt.Enqueue(2)

	if !queueInt.IsFull() {
		t.Errorf("Full was incorrect, got: %t, want: %t.", queueInt.IsFull(), true)

	}
}

func TestIsEmpty(t *testing.T) {

	queueInt := NewQueue[int](2)

	if !queueInt.IsEmpty() {
		t.Errorf("Empty was incorrect, got: %t, want: %t.", queueInt.IsEmpty(), true)

	}
}

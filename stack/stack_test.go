package stack

import (
	"testing"
)

func TestNewStack(t *testing.T) {

	stackInt := NewStack[int](3)

	if len(stackInt.arr) != 3 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", stackInt.size, 3)
	}

	stackStr := NewStack[string](5)

	if len(stackStr.arr) != 5 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", stackStr.size, 5)
	}

}

func TestPush(t *testing.T) {
	stackInt := NewStack[int](3)

	stackInt.Push(21)

	if stackInt.Peek() != 21 {
		t.Errorf("Peek was incorrect, got: %d, want: %d.", stackInt.Peek(), 21)
	}
}

func TestPushMax(t *testing.T) {
	stackInt := NewStack[int](3)

	stackInt.Push(1)
	stackInt.Push(2)
	stackInt.Push(3)

	_, err := stackInt.Push(21)

	if err == nil {
		t.Errorf("Push was incorrect, expected error")
	}

}

func TestPop(t *testing.T) {

	stackInt := NewStack[int](3)

	stackInt.Push(21)
	stackInt.Push(22)

	stackInt.Pop()

	if stackInt.Peek() != 21 {
		t.Errorf("Peek was incorrect, got: %d, want: %d.", stackInt.Peek(), 21)
	}

}

func TestPopEmpty(t *testing.T) {

	stackInt := NewStack[int](3)

	_, err := stackInt.Pop()

	if err == nil {
		t.Errorf("Pop was incorrect, expected error")
	}

}

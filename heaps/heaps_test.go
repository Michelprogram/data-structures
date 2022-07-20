package heaps

import (
	"testing"
)

func TestHeapify(t *testing.T) {

	table := []int{3, 9, 2, 1, 4, 5}

	expected := []int{9, 4, 5, 1, 3, 2}

	heap := NewHeaps()

	for _, value := range table {
		heap.Insert(value)
	}

	if !isEqual(heap.arr, expected) {
		t.Errorf("Heapify was incorrect, got: %v, want: %v.", heap.arr, expected)
	}

}

func TestSwap(t *testing.T) {

	table := []int{3, 9, 2, 1, 4, 5}

	heap := &Heaps{
		arr: table,
	}

	heap.Swap(1, 2)

	if heap.arr[1] != 2 {
		t.Errorf("Heapify was incorrect, got: %v, want: %v.", heap.arr[1], 2)
	}

}

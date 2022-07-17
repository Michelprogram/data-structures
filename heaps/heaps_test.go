package main

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

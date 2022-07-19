package main

type Heaps struct {
	arr  []int
	size int
}

func NewHeaps() *Heaps {

	heaps := Heaps{
		arr:  []int{},
		size: 0,
	}

	return &heaps
}

func (h *Heaps) Insert(element int) {
	h.arr = append(h.arr, element)

	lastKey := len(h.arr) - 1

	h.size++

	h.HeapifyUp(lastKey)
}

//Heapify is the process of creating a heap data structure from a binary tree. Heap Max
func (h *Heaps) HeapifyUp(startIndex int) {

	var parentChild int = h.parentChild(startIndex)

	var minus int = h.arr[parentChild]

	var max int = h.arr[startIndex]

	for minus < max {

		h.Swap(h.parentChild(startIndex), startIndex)

		startIndex = h.parentChild(startIndex)

		parentChild = h.parentChild(startIndex)
		minus = h.arr[parentChild]
		max = h.arr[startIndex]
	}

}

func (h *Heaps) HeapifyDown(startIndex int) {

	lastIndex := len(h.arr) - 1

	l, r := h.leftChild(startIndex), h.rightChild(startIndex)

	compare := 0

	for l <= lastIndex {

		if l == lastIndex {
			compare = l
		} else if h.arr[l] > h.arr[r] {
			compare = l
		} else {
			compare = r
		}

		if h.arr[startIndex] < h.arr[l] {
			h.Swap(startIndex, compare)
			startIndex = compare
			l, r = h.leftChild(startIndex), h.rightChild(startIndex)
		} else {
			return
		}
	}

}

func (h *Heaps) Swap(src, dest int) {
	h.arr[src], h.arr[dest] = h.arr[dest], h.arr[src]
}

func (h *Heaps) leftChild(parentIndex int) int {
	return parentIndex*2 + 1
}

func (h *Heaps) rightChild(parentIndex int) int {
	return parentIndex*2 + 2
}

func (h *Heaps) parentChild(childIndex int) int {
	return (childIndex - 1) / 2
}

func (h *Heaps) Largest() int {
	return h.arr[0]
}

func (h *Heaps) Extract() int {

	if len(h.arr) == 0 {
		return -1
	}

	res := h.arr[0]

	l := len(h.arr) - 1

	h.arr[0] = h.arr[l]

	h.arr = h.arr[:l]

	h.HeapifyDown(h.arr[0])

	return res
}

func isEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

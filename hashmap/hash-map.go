package hashmap

import (
	"fmt"

	"github.com/Michelprogram/data-structures/linkedlist"
)

type Hash struct {
	arr []*Bucket
}

func (h Hash) String() string {

	var res string = ""

	for index := range h.arr {
		res += fmt.Sprintf("At index %d\n", index)

		res += h.arr[index].data.Traversal()

		res += "\n"
	}
	return res
}

func NewHash(size int) *Hash {

	h := &Hash{
		arr: make([]*Bucket, size),
	}

	for i := range h.arr {
		h.arr[i] = NewBucket()
	}

	return h
}

func (h Hash) Insert(key string) {

	indexToInsert := h.Hash(key)

	element := h.arr[indexToInsert]

	if element.data.Search(key) == nil {

		element.data.InsertionBeginning(key)

	}

}

func (h Hash) Search(key string) bool {

	index := h.Hash(key)

	if h.arr[index].data.Search(key) != nil {
		return true
	}

	return false
}

func (h Hash) Delete(key string) bool {

	index := h.Hash(key)


	if h.arr[index].data.Delete(key) != nil {
		return true
	}

	return false
}

func (h Hash) Hash(key string) int {

	var res int = 0

	for _, data := range key {
		res += int(data)
	}
	return res % len(key)
}

type Bucket struct {
	data *linkedlist.LinkedList[string]
}

func NewBucket() *Bucket {

	head := linkedlist.NewNode("-1", nil)

	linklist := linkedlist.NewLinkedList[string]()

	linklist.Head = head

	b := &Bucket{
		data: linklist,
	}

	return b
}

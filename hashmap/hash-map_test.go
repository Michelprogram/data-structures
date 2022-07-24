package hashmap

import (
	"testing"
)

func InitHashMap() *Hash {
	h := NewHash(10)

	h.Insert("darion")

	h.Insert("dorian")

	return h
}

func TestCreateStruct(t *testing.T) {
	h := InitHashMap()

	if len(h.arr) != 10 {
		t.Errorf("Hash array size was incorrect, got: %d, want: %d.", len(h.arr), 10)
	}

	for _, data := range h.arr {
		if data == nil {
			t.Errorf("Hash element was incorrect, got: %v, want: %s.", nil, "Different from nil")
		}
	}

}

func TestHash(t *testing.T) {
	h := InitHashMap()

	if h.Hash("dorian") != 1 {
		t.Errorf("Hash Function was incorrect, got: %d, want: %d.", h.Hash("dorian"), 1)
	}
}

func TestInsert(t *testing.T) {
	h := InitHashMap()

	if h.arr[1].data.Size() != 2 {
		t.Errorf("Hash insert was incorrect, got: %d at index 1, want: %d at index 1.", h.arr[1].data.Size(), 2)
	}
}

func TestSearch(t *testing.T) {
	h := InitHashMap()

	if !h.Search("dorian") {
		t.Errorf("Hash search was incorrect, got: %v, want: %v.", h.Search("dorian"), true)
	}

	if h.Search("eric") {
		t.Errorf("Hash search was incorrect, got: %v, want: %v.", h.Search("eric"), false)
	}
}

func TestDelete(t *testing.T) {
	h := InitHashMap()

	res := h.Delete("darion")

	if h.arr[1].data.Size() != 1 {
		t.Errorf("Hash delete was incorrect, got: %d at index 1, want: %d at index 1.", h.arr[1].data.Size(), 1)
	}

	if !res {
		t.Errorf("Hash delete was incorrect, got: %t, want: %t.", res, true)
	}

}

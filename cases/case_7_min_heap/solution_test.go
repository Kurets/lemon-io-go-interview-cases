package case_7_min_heap

import (
	"testing"
)

func TestInsertAndExtract(t *testing.T) {
	h := &MinHeap{}

	h.Insert(5)
	h.Insert(3)
	h.Insert(8)

	min, ok := h.ExtractMin()
	if !ok || min != 3 {
		t.Errorf("Expected ExtractMin() = 3, got %d, %v", min, ok)
	}

	min, ok = h.ExtractMin()
	if !ok || min != 5 {
		t.Errorf("Expected ExtractMin() = 5, got %d, %v", min, ok)
	}

	min, ok = h.ExtractMin()
	if !ok || min != 8 {
		t.Errorf("Expected ExtractMin() = 8, got %d, %v", min, ok)
	}

	_, ok = h.ExtractMin()
	if ok {
		t.Errorf("Expected ExtractMin() to return false on empty heap, but got true")
	}
}

func TestEmptyHeap(t *testing.T) {
	h := &MinHeap{}

	min, ok := h.ExtractMin()
	if ok {
		t.Errorf("Expected ExtractMin() = false for empty heap, but got true with value %d", min)
	}
}

func TestMaintainsHeapProperty(t *testing.T) {
	h := &MinHeap{}

	values := []int{10, 20, 5, 7, 1, 15, 30}
	for _, v := range values {
		h.Insert(v)
	}

	expected := []int{1, 5, 7, 10, 15, 20, 30}
	for _, exp := range expected {
		min, ok := h.ExtractMin()
		if !ok || min != exp {
			t.Errorf("Expected ExtractMin() = %d, got %d, %v", exp, min, ok)
		}
	}
}

func TestStressTest(t *testing.T) {
	h := &MinHeap{}

	for i := 100; i > 0; i-- {
		h.Insert(i)
	}

	for i := 1; i <= 100; i++ {
		min, ok := h.ExtractMin()
		if !ok || min != i {
			t.Errorf("Expected ExtractMin() = %d, got %d, %v", i, min, ok)
		}
	}

	_, ok := h.ExtractMin()
	if ok {
		t.Errorf("Expected ExtractMin() to return false on empty heap")
	}
}

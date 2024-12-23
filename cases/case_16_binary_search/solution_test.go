package case_16_binary_search

import "testing"

func TestBinarySearchFoundAtBeginning(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	idx := binarySearch(arr, 1)
	if idx != 0 {
		t.Errorf("expected index 0 for target 1, got %d", idx)
	}
}

func TestBinarySearchFoundAtEnd(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	idx := binarySearch(arr, 5)
	if idx != 4 {
		t.Errorf("expected index 4 for target 5, got %d", idx)
	}
}

func TestBinarySearchFoundInMiddle(t *testing.T) {
	arr := []int{2, 4, 6, 8, 10, 12}
	idx := binarySearch(arr, 8)
	if idx != 3 {
		t.Errorf("expected index 3 for target 8, got %d", idx)
	}
}

func TestBinarySearchNotFound(t *testing.T) {
	arr := []int{10, 20, 30, 40, 50}
	idx := binarySearch(arr, 5)
	if idx != -1 {
		t.Errorf("expected -1 for target 5, got %d", idx)
	}
	idx = binarySearch(arr, 60)
	if idx != -1 {
		t.Errorf("expected -1 for target 60, got %d", idx)
	}
	idx = binarySearch(arr, 25)
	if idx != -1 {
		t.Errorf("expected -1 for target 25, got %d", idx)
	}
}

func TestBinarySearchEmptySlice(t *testing.T) {
	arr := []int{}
	idx := binarySearch(arr, 10)
	if idx != -1 {
		t.Errorf("expected -1 for empty slice, got %d", idx)
	}
}

func TestBinarySearchSingleElement(t *testing.T) {
	arr := []int{42}
	idx := binarySearch(arr, 42)
	if idx != 0 {
		t.Errorf("expected index 0 for target 42, got %d", idx)
	}

	idx = binarySearch(arr, 43)
	if idx != -1 {
		t.Errorf("expected -1 for target 43 not in [42], got %d", idx)
	}
}

func TestBinarySearchWithDuplicates(t *testing.T) {
	arr := []int{1, 2, 2, 2, 3}
	idx := binarySearch(arr, 2)
	if idx == -1 {
		t.Error("expected a valid index for target 2, got -1")
	} else {
		if arr[idx] != 2 {
			t.Errorf("expected arr[idx] to be 2, got %d", arr[idx])
		}
	}
}

package case_1_merge_sort

import (
	"reflect"
	"testing"
)

func TestEmptySlice(t *testing.T) {
	input := []int{}
	expected := []int{}
	result := mergeSort(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("mergeSort(%v) = %v; expected %v", input, result, expected)
	}
}

func TestSingleElement(t *testing.T) {
	input := []int{1}
	expected := []int{1}
	result := mergeSort(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("mergeSort(%v) = %v; expected %v", input, result, expected)
	}
}

func TestAlreadySorted(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 3, 4, 5}
	result := mergeSort(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("mergeSort(%v) = %v; expected %v", input, result, expected)
	}
}

func TestReverseOrder(t *testing.T) {
	input := []int{5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5}
	result := mergeSort(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("mergeSort(%v) = %v; expected %v", input, result, expected)
	}
}

func TestWithDuplicates(t *testing.T) {
	input := []int{3, 1, 2, 3, 3, 1}
	expected := []int{1, 1, 2, 3, 3, 3}
	result := mergeSort(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("mergeSort(%v) = %v; expected %v", input, result, expected)
	}
}

func TestUnsorted(t *testing.T) {
	input := []int{5, 2, 9, 1, 5, 6}
	expected := []int{1, 2, 5, 5, 6, 9}
	result := mergeSort(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("mergeSort(%v) = %v; expected %v", input, result, expected)
	}
}

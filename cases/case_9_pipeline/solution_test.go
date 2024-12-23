package case_9_pipeline

import (
	"reflect"
	"testing"
)

func TestGenerator(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	expected := nums

	gen := generator(nums...)
	var result []int
	for n := range gen {
		result = append(result, n)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected generator output %v, got %v", expected, result)
	}
}

func TestFilterEven(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	expected := []int{2, 4}

	gen := generator(nums...)
	even := filterEven(gen)
	var result []int
	for n := range even {
		result = append(result, n)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected filterEven output %v, got %v", expected, result)
	}
}

func TestEmptyInput(t *testing.T) {
	gen := generator()
	var result []int
	for n := range gen {
		result = append(result, n)
	}

	if len(result) != 0 {
		t.Errorf("Expected empty generator output, got %v", result)
	}
}

func TestNoEvens(t *testing.T) {
	nums := []int{1, 3, 5, 7}
	var expected []int

	gen := generator(nums...)
	even := filterEven(gen)
	var result []int
	for n := range even {
		result = append(result, n)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected filterEven output %v, got %v", expected, result)
	}
}

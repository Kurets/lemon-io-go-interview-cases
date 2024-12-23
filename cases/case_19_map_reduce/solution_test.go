package case_19_map_reduce

import (
	"testing"
)

func TestMapReduceBasic(t *testing.T) {
	inputs := []int{1, 2, 3, 4}
	mapper := func(x int) int { return x * x }
	reducer := func(a, b int) int { return a + b }

	result := mapReduce(inputs, mapper, reducer)
	expected := 1 + 4 + 9 + 16

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestMapReduceEmpty(t *testing.T) {
	var inputs []int
	mapper := func(x int) int { return x } // identity
	reducer := func(a, b int) int { return a + b }

	result := mapReduce(inputs, mapper, reducer)
	expected := 0

	if result != expected {
		t.Errorf("Expected %d for empty input, got %d", expected, result)
	}
}

func TestMapReduceSingleElement(t *testing.T) {
	inputs := []int{42}
	mapper := func(x int) int { return x + 1 }
	reducer := func(a, b int) int { return a + b }

	result := mapReduce(inputs, mapper, reducer)
	expected := 43

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestMapReduceDifferentMapperReducer(t *testing.T) {
	inputs := []int{2, 2, 4, 4, 4}
	mapper := func(x int) int { return 1 }
	reducer := func(a, b int) int { return a + b }

	result := mapReduce(inputs, mapper, reducer)
	expected := len(inputs) // 5

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestMapReduceOrderIndependence(t *testing.T) {
	inputs := []int{3, 1, 2}
	mapper := func(x int) int { return x }
	reducer := func(a, b int) int { return a * b }

	result := mapReduce(inputs, mapper, reducer)
	expected := 3 * 1 * 2 // 6

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestMapReduceNegativeNumbers(t *testing.T) {
	inputs := []int{-1, -2, 3}
	mapper := func(x int) int { return x * 2 }
	reducer := func(a, b int) int { return a + b }

	result := mapReduce(inputs, mapper, reducer)
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

package case_24_knapsack

import "testing"

func TestKnapsackEmpty(t *testing.T) {
	var weights []int
	var values []int
	capacity := 10

	result := knapsack(weights, values, capacity)
	expected := 0

	if result != expected {
		t.Errorf("knapsack(empty) = %d, want %d", result, expected)
	}
}

func TestKnapsackZeroCapacity(t *testing.T) {
	weights := []int{1, 2, 3}
	values := []int{6, 10, 12}
	capacity := 0

	result := knapsack(weights, values, capacity)
	expected := 0

	if result != expected {
		t.Errorf("knapsack with zero capacity = %d, want %d", result, expected)
	}
}

func TestKnapsackSingleItemFits(t *testing.T) {
	weights := []int{5}
	values := []int{10}
	capacity := 6

	result := knapsack(weights, values, capacity)
	expected := 10

	if result != expected {
		t.Errorf("knapsack(single item fits) = %d, want %d", result, expected)
	}
}

func TestKnapsackSingleItemDoesNotFit(t *testing.T) {
	weights := []int{5}
	values := []int{10}
	capacity := 4

	result := knapsack(weights, values, capacity)
	expected := 0

	if result != expected {
		t.Errorf("knapsack(single item does not fit) = %d, want %d", result, expected)
	}
}

func TestKnapsackMultipleItems(t *testing.T) {
	weights := []int{1, 3, 4}
	values := []int{15, 50, 60}
	capacity := 4

	result := knapsack(weights, values, capacity)
	expected := 65
	if result != expected {
		t.Errorf("knapsack(multiple items) = %d, want %d", result, expected)
	}
}

func TestKnapsackMultipleItemsAllFit(t *testing.T) {
	weights := []int{2, 3, 1}
	values := []int{4, 5, 3}
	capacity := 10
	result := knapsack(weights, values, capacity)
	expected := 12

	if result != expected {
		t.Errorf("knapsack(all fit) = %d, want %d", result, expected)
	}
}

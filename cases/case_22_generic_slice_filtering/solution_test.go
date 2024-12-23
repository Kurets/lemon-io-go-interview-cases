package case_22_generic_slice_filtering

import (
	"reflect"
	"testing"
)

func TestFilterIntsEven(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	isEven := func(x int) bool { return x%2 == 0 }
	got := Filter(input, isEven)
	want := []int{2, 4, 6}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Filter even numbers failed. Expected %v, got %v", want, got)
	}
}

func TestFilterIntsNoMatches(t *testing.T) {
	input := []int{1, 3, 5, 7}
	isEven := func(x int) bool { return x%2 == 0 }
	got := Filter(input, isEven)
	var want []int

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Filter with no matches failed. Expected %v, got %v", want, got)
	}
}

func TestFilterIntsAllMatch(t *testing.T) {
	input := []int{2, 4, 6, 8}
	isEven := func(x int) bool { return x%2 == 0 }
	got := Filter(input, isEven)
	// all are even, so they all should pass
	if !reflect.DeepEqual(got, input) {
		t.Errorf("Expected %v, got %v", input, got)
	}
}

func TestFilterStrings(t *testing.T) {
	input := []string{"apple", "banana", "pear", "apricot"}
	startsWithA := func(s string) bool { return len(s) > 0 && s[0] == 'a' }
	got := Filter(input, startsWithA)
	want := []string{"apple", "apricot"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Filter strings failed. Expected %v, got %v", want, got)
	}
}

func TestFilterEmptySlice(t *testing.T) {
	var input []int
	isEven := func(x int) bool { return x%2 == 0 }
	got := Filter(input, isEven)
	var want []int

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected empty slice, got %v", got)
	}
}

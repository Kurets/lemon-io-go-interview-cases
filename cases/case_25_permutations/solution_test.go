package case_25_permutations

import (
	"reflect"
	"sort"
	"testing"
)

func sortSliceOfSlices(ss [][]int) {
	for _, s := range ss {
		sort.Ints(s)
	}
	sort.Slice(ss, func(i, j int) bool {
		for k := 0; k < len(ss[i]) && k < len(ss[j]); k++ {
			if ss[i][k] < ss[j][k] {
				return true
			} else if ss[i][k] > ss[j][k] {
				return false
			}
		}
		return len(ss[i]) < len(ss[j])
	})
}

func TestPermutationsEmpty(t *testing.T) {
	var input []int
	got := permutations(input)
	want := [][]int{{}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Permutations of empty slice = %v, want %v", got, want)
	}
}

func TestPermutationsSingle(t *testing.T) {
	input := []int{42}
	got := permutations(input)
	want := [][]int{{42}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Permutations of single element = %v, want %v", got, want)
	}
}

func TestPermutationsTwo(t *testing.T) {
	input := []int{1, 2}
	got := permutations(input)
	sortSliceOfSlices(got)
	want := [][]int{{1, 2}, {2, 1}}
	sortSliceOfSlices(want)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Permutations of [1,2] = %v, want %v", got, want)
	}
}

func TestPermutationsThree(t *testing.T) {
	input := []int{1, 2, 3}
	got := permutations(input)
	sortSliceOfSlices(got)
	want := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}
	sortSliceOfSlices(want)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Permutations of [1,2,3] = %v, want %v", got, want)
	}
}

func TestPermutationsWithDuplicates(t *testing.T) {
	input := []int{1, 1, 2}
	got := permutations(input)
	sortSliceOfSlices(got)
	want := [][]int{
		{1, 1, 2},
		{1, 2, 1},
		{2, 1, 1},
	}
	sortSliceOfSlices(want)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Permutations of [1,1,2] = %v, want %v", got, want)
	}
}

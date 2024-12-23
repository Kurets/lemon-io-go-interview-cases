package case_11_tarjan_algorithm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSingleSCC(t *testing.T) {
	graph := map[int][]int{
		1: {2},
		2: {3},
		3: {1},
	}
	expected := [][]int{{1, 3, 2}}

	result := tarjanSCC(graph)

	if !matchSCCs(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestMultipleSCCs(t *testing.T) {
	graph := map[int][]int{
		1: {2},
		2: {3},
		3: {1},
		4: {5},
		5: {6},
		6: {4},
	}
	expected := [][]int{{1, 3, 2}, {4, 6, 5}}

	result := tarjanSCC(graph)

	if !matchSCCs(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestNoEdges(t *testing.T) {
	graph := map[int][]int{
		1: {},
		2: {},
		3: {},
	}
	expected := [][]int{{1}, {2}, {3}}

	result := tarjanSCC(graph)

	if !matchSCCs(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestOneLargeCycle(t *testing.T) {
	graph := map[int][]int{
		1: {2},
		2: {3},
		3: {4},
		4: {1},
	}
	expected := [][]int{{1, 4, 3, 2}}

	result := tarjanSCC(graph)

	if !matchSCCs(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDisconnectedGraph(t *testing.T) {
	graph := map[int][]int{
		1: {2},
		2: {1},
		3: {4},
		4: {},
		5: {},
	}
	expected := [][]int{{1, 2}, {3}, {4}, {5}}

	result := tarjanSCC(graph)

	if !matchSCCs(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestEmptyGraph(t *testing.T) {
	graph := map[int][]int{}
	expected := [][]int{}

	result := tarjanSCC(graph)

	if !matchSCCs(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func matchSCCs(result, expected [][]int) bool {
	resultMap := make(map[string]bool)
	expectedMap := make(map[string]bool)

	for _, scc := range result {
		resultMap[hashSCC(scc)] = true
	}
	for _, scc := range expected {
		expectedMap[hashSCC(scc)] = true
	}

	return reflect.DeepEqual(resultMap, expectedMap)
}

func hashSCC(scc []int) string {
	m := make(map[int]bool)
	for _, v := range scc {
		m[v] = true
	}
	return fmt.Sprintf("%v", m)
}

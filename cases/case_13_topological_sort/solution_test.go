package case_13_topological_sort

import (
	"testing"
)

func isValidTopologicalOrder(graph map[int][]int, order []int) bool {
	pos := make(map[int]int)
	for i, node := range order {
		pos[node] = i
	}

	for u := range graph {
		if _, ok := pos[u]; !ok {
			return false
		}
	}

	for u, neighbors := range graph {
		for _, v := range neighbors {
			if pos[u] > pos[v] {
				return false
			}
		}
	}
	return true
}

func TestBasicDAG(t *testing.T) {
	graph := map[int][]int{
		0: {1, 2},
		1: {3},
		2: {3},
		3: {},
	}

	order := topologicalSort(graph)
	if !isValidTopologicalOrder(graph, order) {
		t.Errorf("topologicalSort returned an invalid order: %v", order)
	}
}

func TestMultipleValidOrders(t *testing.T) {
	graph := map[int][]int{
		0: {2},
		1: {2},
		2: {},
	}

	order := topologicalSort(graph)
	if !isValidTopologicalOrder(graph, order) {
		t.Errorf("topologicalSort returned an invalid order for a graph with multiple valid answers: %v", order)
	}
}

func TestSingleNode(t *testing.T) {
	graph := map[int][]int{
		0: {},
	}
	order := topologicalSort(graph)
	if !isValidTopologicalOrder(graph, order) {
		t.Errorf("Invalid order for single node graph: %v", order)
	}
	if len(order) != 1 || order[0] != 0 {
		t.Errorf("Expected [0], got %v", order)
	}
}

func TestNoEdges(t *testing.T) {
	graph := map[int][]int{
		0: {},
		1: {},
		2: {},
	}
	order := topologicalSort(graph)
	if len(order) != 3 {
		t.Errorf("Expected all nodes to appear in the order, got %v", order)
	}

	if !isValidTopologicalOrder(graph, order) {
		t.Errorf("Order %v should be valid since there are no edges", order)
	}
}

func TestEmptyGraph(t *testing.T) {
	graph := map[int][]int{}
	order := topologicalSort(graph)

	if len(order) != 0 {
		t.Errorf("Expected empty order for empty graph, got %v", order)
	}
}

package case_12_dijkstra_algorithm

import (
	"math"
	"reflect"
	"testing"
)

func TestBasicGraph(t *testing.T) {
	graph := map[int][]struct{ to, weight int }{
		0: {{to: 1, weight: 4}, {to: 2, weight: 1}},
		1: {{to: 3, weight: 1}},
		2: {{to: 1, weight: 2}, {to: 3, weight: 5}},
		3: {},
	}

	dist := dijkstra(graph, 0)
	expected := map[int]int{
		0: 0,
		1: 3,
		2: 1,
		3: 4,
	}

	if !reflect.DeepEqual(dist, expected) {
		t.Errorf("expected distances %v, got %v", expected, dist)
	}
}

func TestUnreachableNodes(t *testing.T) {
	// Here, nodes 4 and 5 are not reachable from 0.
	graph := map[int][]struct{ to, weight int }{
		0: {{to: 1, weight: 2}},
		1: {{to: 2, weight: 2}},
		2: {{to: 3, weight: 2}},
		3: {},
		4: {{to: 5, weight: 1}},
		5: {},
	}

	dist := dijkstra(graph, 0)
	// We know from 0: dist[0]=0, dist[1]=2, dist[2]=4, dist[3]=6
	// Nodes 4 and 5 are unreachable, so their distance should remain MaxInt
	expected := map[int]int{
		0: 0,
		1: 2,
		2: 4,
		3: 6,
		4: math.MaxInt,
		5: math.MaxInt,
	}
	if !reflect.DeepEqual(dist, expected) {
		t.Errorf("expected distances %v, got %v", expected, dist)
	}
}

func TestSingleNode(t *testing.T) {
	graph := map[int][]struct{ to, weight int }{
		0: {},
	}

	dist := dijkstra(graph, 0)
	expected := map[int]int{
		0: 0,
	}
	if !reflect.DeepEqual(dist, expected) {
		t.Errorf("expected %v, got %v", expected, dist)
	}
}

func TestNoEdges(t *testing.T) {
	// Multiple nodes but no edges.
	graph := map[int][]struct{ to, weight int }{
		0: {},
		1: {},
		2: {},
	}

	dist := dijkstra(graph, 0)
	// Only node 0 is known and is start, so dist[0]=0, others remain MaxInt
	expected := map[int]int{
		0: 0,
		1: math.MaxInt,
		2: math.MaxInt,
	}

	if !reflect.DeepEqual(dist, expected) {
		t.Errorf("expected %v, got %v", expected, dist)
	}
}

func TestComplexGraph(t *testing.T) {
	// A more complex graph:
	//   0 ->1(2), 0->2(4)
	//   1->2(1), 1->3(7)
	//   2->3(1), 2->4(3)
	//   3->4(2)
	// Shortest paths from 0:
	//  0 to 0 = 0
	//  0 to 1 = 2 (direct)
	//  0 to 2 = 3 (0->1->2)
	//  0 to 3 = 4 (0->1->2->3 or 0->2->3)
	//  0 to 4 = 6 (0->1->2->4 or 0->2->4)
	graph := map[int][]struct{ to, weight int }{
		0: {{to: 1, weight: 2}, {to: 2, weight: 4}},
		1: {{to: 2, weight: 1}, {to: 3, weight: 7}},
		2: {{to: 3, weight: 1}, {to: 4, weight: 3}},
		3: {{to: 4, weight: 2}},
		4: {},
	}

	dist := dijkstra(graph, 0)
	expected := map[int]int{
		0: 0,
		1: 2,
		2: 3,
		3: 4,
		4: 6,
	}

	if !reflect.DeepEqual(dist, expected) {
		t.Errorf("expected %v, got %v", expected, dist)
	}
}

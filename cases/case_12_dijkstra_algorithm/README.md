# Dijkstra's Algorithm for Shortest Paths

## Task Description

Implement **Dijkstra's algorithm** to find the shortest distances from a given start node to all other nodes in a
weighted, directed graph. The graph is represented as an adjacency list of the form `map[int][]struct{to, weight int}`.

### Function Signature

```go
package main

func dijkstra(graph map[int][]struct{ to, weight int }, start int) map[int]int {
	// Your code here
	return nil
}
```

### Requirements

1. **Input**:
    - `graph`: A directed graph represented as an adjacency list where:
        - Keys are node IDs.
        - Values are slices of structs containing:
            - `to`: The destination node.
            - `weight`: The weight of the edge.
    - `start`: The node from which to calculate shortest distances.

2. **Output**:
    - A map where:
        - Keys are node IDs.
        - Values are the shortest distances from the `start` node to the respective node.
    - If a node is not reachable from the `start`, its distance should be `math.MaxInt`.

3. **Behavior**:
    - Use a **priority queue** (min-heap) to always process the node with the smallest known distance.
    - Relax edges to update the shortest known distances.
    - Return the computed shortest distances for all nodes.

4. **Complexity**:
    - Time Complexity: **O((V + E) log V)**, where `V` is the number of nodes and `E` is the number of edges.

### Example Usage

```go
package main

import "fmt"

func main() {
	graph := map[int][]struct{ to, weight int }{
		0: {{to: 1, weight: 4}, {to: 2, weight: 1}},
		1: {{to: 3, weight: 1}},
		2: {{to: 1, weight: 2}, {to: 3, weight: 5}},
		3: {},
	}
	start := 0
	dist := dijkstra(graph, start)

	fmt.Println(dist)
	// Output:
	// map[0:0 1:3 2:1 3:4]

}
```

### Notes

- **Algorithm Details**:
    - Use a priority queue to manage nodes and their current shortest distances.
    - Initialize all distances to `math.MaxInt` (infinity), except the `start` node which is set to `0`.
    - Process each node:
        - For each neighbor, calculate the distance through the current node.
        - If the calculated distance is shorter than the known distance, update it and push the neighbor into the
          priority queue.

- **Edge Cases**:
    - A graph with no nodes or edges.
    - Disconnected graphs where some nodes are unreachable.
    - A single node graph.

### Hints

- **Priority Queue**:
    - Use a min-heap (`container/heap`) to efficiently retrieve the next node with the smallest distance.

- **Relaxation**:
    - For each edge `(u, v)` with weight `w`:
        - If `dist[u] + w < dist[v]`, update `dist[v]` to `dist[u] + w`.

- **Initialization**:
    - Set the distance of all nodes to `math.MaxInt`.
    - Set the distance of the `start` node to `0`.

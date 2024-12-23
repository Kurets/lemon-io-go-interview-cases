# Topological Sort for Directed Acyclic Graph (DAG)

## Task Description

Implement a function to perform **topological sort** on a directed acyclic graph (DAG). The function should return one
possible ordering of nodes such that for every directed edge `(u, v)`, node `u` appears before node `v` in the ordering.

### Function Signature

```go
package main

func topologicalSort(graph map[int][]int) []int {
	// Your code here
	return nil
}
```

### Requirements

1. **Input**:
    - `graph`: A directed acyclic graph represented as an adjacency list:
        - Keys are node IDs.
        - Values are slices of integers representing neighboring nodes.

2. **Output**:
    - A slice of integers representing one valid topological ordering of the nodes.

3. **Behavior**:
    - Compute the **in-degree** (number of incoming edges) for each node.
    - Initialize a queue with all nodes that have an in-degree of 0.
    - Process nodes in the queue:
        - Append the node to the order.
        - Decrement the in-degree of its neighbors.
        - Add neighbors with an in-degree of 0 to the queue.
    - Return the order once all nodes are processed.

4. **Complexity**:
    - Time Complexity: **O(V + E)**, where `V` is the number of nodes and `E` is the number of edges.

### Example Usage

```go
package main

import "fmt"

func main() {
	graph := map[int][]int{
		0: {1, 2},
		1: {3},
		2: {3},
		3: {},
	}

	order := topologicalSort(graph)
	fmt.Println(order)
	// Output: [0, 1, 2, 3] or [0, 2, 1, 3] (any valid topological order)

}
```

### Notes

- **Topological Sort**:
    - Only applicable to directed acyclic graphs (DAGs).
    - A valid ordering ensures that for every edge `(u, v)`, `u` appears before `v`.

- **Algorithm Details**:
    - Compute in-degrees for all nodes.
    - Use a queue to process nodes with in-degree 0.
    - Append nodes to the order and update in-degrees of their neighbors.

### Edge Cases

1. **Empty Graph**:
    - Input graph has no nodes (`graph == nil`).
    - Output should be an empty slice.

2. **Disconnected Graph**:
    - Graph with multiple nodes but no edges.
    - Any permutation of nodes is a valid topological order.

3. **Single Node**:
    - Graph contains a single node with no edges.
    - The order should contain just that node.

4. **Multiple Valid Orders**:
    - Graph with multiple valid topological orders.
    - The function can return any valid order.

### Hints

- **In-degree Calculation**:
    - Traverse the adjacency list to count the incoming edges for each node.

- **Queue Initialization**:
    - Add nodes with an in-degree of 0 to the queue at the start.

- **Order Construction**:
    - Use a queue to process nodes in in-degree order.
    - For each node, decrement the in-degree of its neighbors and enqueue any neighbor whose in-degree becomes 0.

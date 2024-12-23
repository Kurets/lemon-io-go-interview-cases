# Tarjan's Algorithm for Strongly Connected Components (SCCs)

## Task Description

Implement **Tarjan's Algorithm** to find all strongly connected components (SCCs) in a directed graph. The graph is
represented as an adjacency list: `map[int][]int`.

### Function Signature

```go
package main

func tarjanSCC(graph map[int][]int) [][]int {
	// Your code here
	return nil
}
```

### Requirements

1. **Input**:
    - `graph`: A directed graph represented as an adjacency list where the keys are node IDs and the values are slices
      of neighboring nodes.

2. **Output**:
    - A slice of slices, where each inner slice represents an SCC containing node IDs.
    - SCCs can be in any order, but the nodes within each SCC must be grouped together.

3. **Behavior**:
    - Use **Tarjan's Algorithm** to find SCCs:
        - Perform Depth-First Search (DFS).
        - Maintain `index` and `low-link` values for each node.
        - Use a stack to track nodes currently in the DFS stack.
        - Identify SCCs when the `low-link` value equals the `index` value for a node.

4. **Complexity**:
    - Time Complexity: **O(V + E)**, where `V` is the number of nodes and `E` is the number of edges.

### Example Usage

```go
package main

import "fmt"

func main() {
	graph := map[int][]int{
		1: {2},
		2: {3},
		3: {1, 4},
		4: {5},
		5: {6, 7},
		6: {4},
		7: {},
	}

	fmt.Println(tarjanSCC(graph))
	// Output (order of SCCs may vary): [[1, 3, 2], [4, 6], [5], [7]]

}
```

### Notes

- **Strongly Connected Components**:
    - An SCC is a maximal subgraph where every node is reachable from every other node in the subgraph.
    - Example:
        - For the graph `1 -> 2 -> 3 -> 1`, the SCC is `[1, 2, 3]`.

- **Algorithm Details**:
    - Maintain:
        - `index`: Tracks the order of discovery for each node.
        - `low-link`: The smallest index reachable from the node.
    - Use a stack to manage nodes in the current DFS path.

### Edge Cases

1. An empty graph (`graph == nil`).
2. A graph with no edges (`graph[node] == []` for all nodes).
3. Disconnected graphs.
4. Cycles of varying sizes.
5. Single-node SCCs (e.g., nodes with no edges).

### Hints

- **Initialization**:
    - Use maps to track `index` and `low-link` values for nodes.
    - Use a boolean map to track whether a node is currently in the DFS stack.

- **Algorithm Steps**:
    1. Start DFS from an unvisited node.
    2. Set its `index` and `low-link` values to the current index.
    3. Push the node onto the stack and mark it as in the stack.
    4. For each neighbor:
        - If unvisited, recursively call DFS and update the node's `low-link` value.
        - If in the stack, update the `low-link` value based on the neighbor's `index`.
    5. If `low-link[node] == index[node]`, extract an SCC by popping nodes from the stack until the current node is
       reached.

- **Output**:
    - Collect SCCs during the DFS process and return them at the end.

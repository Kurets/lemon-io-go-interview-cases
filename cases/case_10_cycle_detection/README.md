# Cycle Detection in Singly Linked List

## Task Description

Implement a function to detect if a singly linked list has a cycle. A cycle exists if a node’s `Next` pointer points
back to a previous node, creating a loop.

### Function Signature

```go
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	// Your code here
	return false
}

```

### Requirements

1. **Input**:
    - `head`: The head of a singly linked list (`*ListNode`).

2. **Output**:
    - Returns `true` if a cycle is detected, otherwise `false`.

3. **Behavior**:
    - Use **Floyd's Cycle Detection Algorithm** (Tortoise and Hare):
        - Use two pointers: `slow` moves one step at a time, and `fast` moves two steps.
        - If `slow` and `fast` meet, a cycle exists.
        - If `fast` or `fast.Next` becomes `nil`, the list does not have a cycle.

### Example Usage

```go
package main

import "fmt"

func main() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n1.Next = n2
	n2.Next = n3

	// Test without cycle
	fmt.Println(hasCycle(n1)) // Output: false

	// Create a cycle
	n3.Next = n1
	fmt.Println(hasCycle(n1)) // Output: true

}
```

### Notes

- **Cycle**:
    - A cycle exists if a node points back to an earlier node in the list.
    - Example: `1 -> 2 -> 3 -> 4 -> 2` (cycle starts at node with value 2).

- **Algorithm Details**:
    - Use two pointers (`slow` and `fast`) to traverse the list.
    - If `slow` equals `fast`, a cycle is detected.
    - If the list ends (`fast == nil` or `fast.Next == nil`), there is no cycle.

### Edge Cases

1. An empty list (`head == nil`).
2. A single node with no cycle.
3. A single node pointing to itself (cycle).
4. Multiple nodes with and without cycles.
5. Cycles at various positions in the list (beginning, middle, or end).

### Hints

- **Detecting Cycle**:
    - Use a two-pointer technique to traverse the list at different speeds.
    - If they meet, the list contains a cycle.

- **Avoid Infinite Loops**:
    - Use the fast pointer's ability to skip nodes to terminate the loop if no cycle exists.

# Min-Heap

## Task Description

Implement a **min-heap** with the following methods:

1. **Insert(val int)**:
    - Adds an integer to the heap while maintaining the heap property:
        - The smallest element is always at the root.
        - Each parent node is smaller than its child nodes.

2. **ExtractMin() (int, bool)**:
    - Removes and returns the smallest element (root) from the heap.
    - Returns `false` if the heap is empty.

### Function Signatures

```go
package main

type MinHeap struct {
	// Your code here
}

func (h *MinHeap) Insert(val int) {
	// Your code here
}

func (h *MinHeap) ExtractMin() (int, bool) {
	// Your code here
	return 0, false
}
```

### Requirements

1. **Heap Property**:
    - The heap should maintain the min-heap property:
        - The smallest element is always at the root (`h.data[0]`).
        - The children of index `i` are at indices `2i+1` and `2i+2`.
        - The parent of index `i` is at `(i-1)/2`.

2. **Insert Method**:
    - Add the new element to the end of the heap.
    - Restore the heap property by "bubbling up" the element.

3. **ExtractMin Method**:
    - Swap the root with the last element, then remove the last element.
    - Restore the heap property by "bubbling down" the root element.

4. **Edge Cases**:
    - Handle extraction from an empty heap.
    - Ensure the heap property is maintained after multiple insertions and extractions.

### Example Usage

```go
package main

import "fmt"

func main() {
	h := &MinHeap{}

	h.Insert(5)
	h.Insert(3)
	h.Insert(8)

	fmt.Println(h.ExtractMin()) // Output: 3, true
	fmt.Println(h.ExtractMin()) // Output: 5, true
	fmt.Println(h.ExtractMin()) // Output: 8, true
	fmt.Println(h.ExtractMin()) // Output: 0, false (heap is empty)

}
```

### Notes

- Use a slice (`[]int`) as the underlying data structure for the heap.
- Implement helper functions to "bubble up" and "bubble down" elements as needed.
- Ensure that all operations maintain the min-heap property efficiently.

### Hints

- **Bubble Up**:
    - Starting from the last inserted element, compare it with its parent.
    - Swap the element with its parent if it is smaller, and continue until the heap property is restored.

- **Bubble Down**:
    - Starting from the root, compare it with its children.
    - Swap it with the smaller child if it violates the heap property, and continue until the property is restored.

- **Index Calculations**:
    - Parent index: `(i-1)/2`
    - Left child index: `2i+1`
    - Right child index: `2i+2`

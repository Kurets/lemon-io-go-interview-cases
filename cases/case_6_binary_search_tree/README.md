# Binary Search Tree (BST)

## Task Description

Implement a **binary search tree** with the following methods:

1. **Insert(val int)**:
    - Inserts a value into the binary search tree.
    - Ensures the tree maintains the binary search property:
        - Values in the left subtree are less than the current node.
        - Values in the right subtree are greater than the current node.
        - Duplicate values should not be added.

2. **Search(val int) bool**:
    - Searches for a value in the binary search tree.
    - Returns `true` if the value is found, otherwise returns `false`.

### Function Signatures

```go
package main

type BST struct {
	val   int
	left  *BST
	right *BST
}

func (b *BST) Insert(val int) {
	// Your code here
}

func (b *BST) Search(val int) bool {
	// Your code here
	return false
}
```

### Requirements

1. **Insert Method**:
    - Adds the value in the correct position based on the binary search property.
    - Avoids adding duplicate values.

2. **Search Method**:
    - Recursively checks the left or right subtree depending on the value being searched.
    - Returns `false` if a `nil` node is reached, indicating the value is not in the tree.

3. **Edge Cases**:
    - Inserting into an empty tree.
    - Searching in an empty tree.
    - Duplicate insertions should not modify the tree.

### Example Usage

```go
package main

import "fmt"

func main() {
	var tree BST

	// Insert values into the BST
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)

	// Search for values
	fmt.Println(tree.Search(3))  // Output: true
	fmt.Println(tree.Search(10)) // Output: false

	// Insert more values
	tree.Insert(4)
	tree.Insert(6)

	// Validate structure manually or via tests
}
```

### Notes

- Ensure the binary search property is maintained after every insertion.
- Avoid duplicate nodes when inserting a value that already exists in the tree.
- Use recursion to simplify the implementation of both `Insert` and `Search`.

### Hints

- **Insert Method**:
    - Check if the value is less than the current node's value to decide whether to insert into the left subtree.
    - Check if the value is greater to decide whether to insert into the right subtree.
    - If the node is `nil`, create a new node with the given value.

- **Search Method**:
    - Use recursion to traverse the tree.
    - Return `true` when the current node's value matches the search value.
    - Return `false` if a `nil` node is reached.

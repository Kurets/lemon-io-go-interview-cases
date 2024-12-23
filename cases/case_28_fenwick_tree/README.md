# Fenwick (Binary Indexed) Tree

## Task Description

You need to implement a **Fenwick Tree** (also called a **Binary Indexed Tree**, or BIT) for prefix sums:

```go
package main

type Fenwick struct {
	// Your code here
}

func NewFenwick(size int) *Fenwick {
	// Your code here
	return nil
}

func (f *Fenwick) Update(i, delta int) {
	// Your code here
}

func (f *Fenwick) PrefixSum(i int) int {
	// Your code here
	return 0
}
```

### Requirements

1. **Fenwick Tree Structure**
    - Backed by an array (often 1-indexed) where each element covers a range of indices determined by its last set bit.
    - Typically `tree[i]` stores the partial sum of elements from `i - (i & -i) + 1` through `i` inclusive.

2. **`Update(i, delta int)`**
    - Increases the value at index `i` by `delta`.
    - Then propagate that update to all relevant Fenwick tree nodes.

3. **`PrefixSum(i int) int`**
    - Returns the sum of elements from index `1` up to index `i` (inclusive).
    - Traverses the Fenwick tree array in a way that accumulates partial sums.

4. **1-Based Indexing**
    - Often Fenwick trees are implemented using 1-based indexing. For a tree of size `n`, valid indices are `1`..`n`.
    - If your input is 0-based, you may need to adjust by +1.

5. **Complexities**
    - Both `Update` and `PrefixSum` should work in O(log n) time.

6. **Edge Cases**
    - **Size=1**: minimal tree.
    - **Index out of bounds**: Usually the user of the tree is expected to remain in `[1..size]`.
    - **Negative or zero updates** are allowed and should function normally.

### Example Usage

```go
package main

import "fmt"

func main() {
	fenw := NewFenwick(5)
	fenw.Update(1, 2)
	fenw.Update(3, 5)
	// PrefixSum(3) => 2 (index 1) + 5 (index 3) = 7
	fmt.Println("PrefixSum(3) =", fenw.PrefixSum(3))
}
```

### Hints

- For a Fenwick tree, the typical code patterns are:
    - `Update(i, delta)`:
    ```go
    	for i <= size {
    		tree[i] += delta
    		i += i & (-i)
    	}
    ```
    - `PrefixSum(i)`:  
     ```go
         result := 0
         for i > 0 {
             result += tree[i]
             i -= i & (-i)
         }
         return result
     ```
- Make sure you handle 1-based indexing consistently.  

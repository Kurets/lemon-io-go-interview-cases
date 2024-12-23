# Merge Sort

## Task Description

Implement a **merge sort** function that takes a slice of integers and returns a new sorted slice using the merge sort
algorithm.

```go
package main

func mergeSort(nums []int) []int {
   // Your code here
   return nil
}

```

### Requirements

1. **Input**
    - A slice of integers, e.g., `[]int{5, 2, 9, 1}`.

2. **Output**
    - A new slice that contains the integers in sorted order.

3. **Algorithm Details**
    - **Divide and Conquer**:
        - Recursively divide the input slice into two halves until each half contains one or no elements.
        - Merge the sorted halves to produce the final sorted slice.
    - The merging step should ensure elements are in ascending order.

4. **Edge Cases**
    - **Empty Slice**: Return an empty slice.
    - **Single Element**: Return the same slice.
    - **Already Sorted**: Ensure the output is identical to the input.

### Example Usage

```go
package main

import (
   "fmt"
)

func main() {
   nums := []int{5, 2, 9, 1, 6, 3}
   sorted := mergeSort(nums)
   fmt.Println(sorted) // [1, 2, 3, 5, 6, 9]
}
```

### Hints

- Use a helper function to perform the **merge** step, where two sorted slices are combined into one sorted slice.
- The base case for the recursion is when the input slice has one or no elements.
- Ensure that your implementation works for slices of varying lengths, including empty slices.

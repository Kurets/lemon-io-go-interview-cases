# Binary Search

## Task Description

Implement a function that performs **binary search** to find the index of a target integer in a sorted slice. If the
target is found, return its index; otherwise, return `-1`.

### Function Signature

```go
package main

func binarySearch(arr []int, target int) int {
	// Your code here
	return -1
}
```

### Requirements

1. **Input**:
    - `arr`: A sorted slice of integers.
    - `target`: An integer to search for in the slice.

2. **Output**:
    - Return the index of `target` if found.
    - Return `-1` if `target` is not in the slice.

3. **Behavior**:
    - Use a binary search algorithm:
        1. Start with two pointers, `low` (0) and `high` (len(arr) - 1).
        2. Calculate `mid = (low + high) / 2`.
        3. Compare `arr[mid]` with `target`:
            - If equal, return `mid`.
            - If `arr[mid] < target`, move `low` to `mid + 1`.
            - If `arr[mid] > target`, move `high` to `mid - 1`.
        4. Repeat until `low > high`.

4. **Complexity**:
    - Time Complexity: **O(log n)**.
    - Space Complexity: **O(1)** (iterative implementation).

### Example Usage

```go
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(binarySearch(arr, 3)) // Output: 2
	fmt.Println(binarySearch(arr, 6)) // Output: -1
}
```

### Notes

- **Binary Search**:
    - Efficient for searching in a sorted slice.
    - Works by repeatedly dividing the search space in half.

- **Edge Cases**:
    - Empty slice.
    - Target less than the smallest element.
    - Target greater than the largest element.
    - Target not in the slice but within the range of elements.

### Edge Cases

1. **Empty Slice**:
    - Input slice is empty (`arr == nil` or `len(arr) == 0`).
    - Output should be `-1`.

2. **Single Element**:
    - Slice has one element:
        - If the target matches the element, return its index (0).
        - Otherwise, return `-1`.

3. **Duplicates**:
    - If the slice contains duplicate elements, any valid index of the target is acceptable.

4. **Not Found**:
    - Target is outside the range of the smallest and largest elements.
    - Target is within the range but not in the slice.

5. **Boundary Conditions**:
    - Target matches the first or last element in the slice.

### Hints

- **Midpoint Calculation**:
    - Use `mid := low + (high-low)/2` to avoid potential overflow in some environments.

- **Loop Termination**:
    - Stop the loop when `low > high`.

- **Testing for Duplicates**:
    - Ensure the returned index corresponds to a valid occurrence of the target, even if there are duplicates.

### Example Workflow

1. Start with the entire range of the slice (`low = 0`, `high = len(arr) - 1`).
2. Calculate the midpoint and compare:
    - If `arr[mid] == target`, return `mid`.
    - If `arr[mid] < target`, discard the left half by updating `low`.
    - If `arr[mid] > target`, discard the right half by updating `high`.
3. Repeat until the range is invalid (`low > high`).

### Performance Considerations

- Binary search is optimal for sorted slices.
- Ensure correctness for edge cases like empty slices and single-element slices.

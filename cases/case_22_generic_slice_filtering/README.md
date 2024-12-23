# Generic Slice Filtering

## Task Description

Using Go generics, implement a function `Filter` that:

``` go 
package main

func Filter[T any](slice []T, pred func(T) bool) []T {
	// Your code here
	return nil
}
```

### Requirements

1. **Input**
    - A slice of any type `T`.
    - A predicate function `func(T) bool`.

2. **Output**
    - A new slice containing **only** the elements for which `pred` returns `true`.
    - Retain the order of elements that satisfy the predicate.

3. **Constraints & Behavior**
    - Do **not** modify the original slice in place.
    - Return an empty slice if no elements match or if the original slice is empty.
    - Maintain type safety using Go generics.

4. **Edge Cases**
    - **Empty slice** => returns empty slice.
    - **No elements match** => returns empty slice.
    - **All elements match** => returns a new slice identical to the original.

### Example Usage

``` go
package main

import (
"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	isOdd := func(x int) bool {
		return x%2 != 0
	}
	oddNums := Filter(nums, isOdd)
	fmt.Println(oddNums) // [1, 3, 5]
}
```

### Hints

- Use a simple loop to accumulate elements for which `pred(element)` is `true`.
- Since `T` is a generic type, your function should work equally for `[]int`, `[]string`, or other slice types.
- Return a new slice; do **not** mutate or reorder the original slice.

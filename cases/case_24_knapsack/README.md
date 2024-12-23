# 0/1 Knapsack with Dynamic Programming

## Task Description

Implement a **0/1 knapsack** function:

```go
package main

func knapsack(weights, values []int, capacity int) int {
	// Your code here
	return 0
}
```

### Requirements

1. **Inputs**
    - `weights`: Slice of item weights (each corresponds to an item).
    - `values`: Slice of item values, same length as `weights`.
    - `capacity`: Maximum weight capacity of the knapsack.

2. **Output**
    - Integer representing the **maximum total value** that can be achieved without exceeding `capacity`.

3. **Constraints**
    - This is the classic **0/1 knapsack** problem, meaning each item may be **taken or not** taken (no partial item).

4. **Approach**
    - Use **dynamic programming** to compute the best solution:
        - Typically a 2D table approach: `dp[i][w]` storing the best possible value using items up to `i` with capacity `w`.
        - Or a 1D rolling array optimization is also acceptable.

5. **Edge Cases**
    - **Empty** `weights`/`values`: The result is 0.
    - **Zero capacity**: No items can be taken, so result is 0.
    - **All items fit**: Then result is the sum of all `values`.
    - **Large capacity** with big or small items.
    - **Weights** and **values** must have the same length (may assume the caller ensures this).

### Example Usage

```go
package main

import "fmt"

func main() {
	weights := []int{1, 3, 4}
	values := []int{15, 50, 60}
	capacity := 4

	maxVal := knapsack(weights, values, capacity)
	fmt.Println("Max knapsack value:", maxVal) // Should print 65 for items (1 & 2)
}

```

### Hints

- A common DP solution is:
    1. Initialize a table `dp[len(weights)+1][capacity+1]` to 0.
    2. For each item `i`, for each possible weight `w` from 0..capacity:
        - Either you **don’t take** item `i`, so `dp[i][w] = dp[i-1][w]`.
        - Or you **take** item `i` (if it fits), so `dp[i][w] = dp[i-1][w - weights[i]] + values[i]`.
        - Then `dp[i][w] = max(of those two)`.
    3. Return `dp[len(weights)][capacity]`.
- You can also optimize space usage to 1D, iterating weights from high to low to avoid overwriting needed subproblems.  

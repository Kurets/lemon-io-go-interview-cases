# Generating All Permutations

## Task Description

Implement a function that generates **all permutations** of a given slice of integers:

``` go
package main

func permutations(nums []int) [][]int {
// Your code here
return nil
}
```

### Requirements

1. **Input**
    - A slice of integers (e.g., `[1,2,3]`).

2. **Output**
    - All unique permutations of the input slice, each permutation represented as a slice of integers.
    - The function returns a slice of these permutations (a `[][]int`).

3. **Behavior**
    - Permutations are the distinct arrangements of the given elements.
    - For `n` distinct elements, there are `n!` permutations.
    - If duplicates exist in `nums`, permutations containing the same elements in the same order should not be repeated multiple times.

4. **Edge Cases**
    - **Empty slice** => The only “permutation” is the empty slice `[]`. Return `[[]]`.
    - **Single element** => Return `[ [elem] ]`.
    - **Duplicates** => Must avoid generating the same permutation multiple times.

5. **Example**
    - Input: `[1, 2, 3]`
        - Output: `[ [1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1] ]` (order can vary).

### Example Usage

``` go
package main

import "fmt"

func main() {
data := []int{1, 2, 3}
perms := permutations(data)
fmt.Println("Permutations of", data, "are:", perms)
}
```

### Hints

- You can generate permutations by:
    1. **Recursive approach**:
        - Swap elements in the slice, recurse, swap back.
        - Keep track of used elements or do a backtracking approach.
    2. **Use a helper** that builds the result by picking each element in turn, recursing on the remainder.
    3. **Handle duplicates** by sorting the input and skipping repeated elements in the recursion or by using a `map`/`set` approach.

- Make sure to **avoid** generating the same permutation multiple times if the input contains duplicates.
- The test includes a case with duplicates (`[1, 1, 2]`). You must ensure your solution returns each distinct arrangement exactly once.  


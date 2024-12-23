# Map-Reduce Pattern

## Task Description

You must implement a **map-reduce** pattern in Go:

```go
func mapReduce(inputs []int, mapper func(int) int, reducer func(int, int) int) int {
    // Your code here
    return 0
}
```

### Requirements

- **Inputs**: A slice of integers (`inputs`).
- **Mapper**: A function `func(int) int` that transforms each integer into an **intermediate** integer result.
    - Example: squaring the number, incrementing by 1, turning each integer into “1” to count occurrences, etc.
- **Reducer**: A function `func(int, int) int` that **aggregates** all intermediate results into a single integer.
    - Example: sum, product, minimum, maximum, etc.
- **Behavior**:
    - You must apply `mapper` to **each** element of `inputs`.
    - Then you must **reduce** all mapper outputs into one final integer result using `reducer`.
    - Return this final integer.
- **Edge Cases**:
    - **Empty inputs**: The behavior might default to 0 or some neutral value, depending on your design.
    - **Single-element** input: The result is simply `mapper(singleValue)`.
    - **Large input**: Implementation should handle any size of input array without issues.

### Example Usage

```go
func main() {
	inputs := []int{1,2,3,4}
	mapper := func(x int) int { return x*x }         // square
	reducer := func(a int, b int) int { return a+b } // sum

	result := mapReduce(inputs, mapper, reducer)
	fmt.Println("Sum of squares:", result)
	// Expected: 1+4+9+16 = 30
}
```

### Hints

- Think of it like two steps:
    1. **Map**: transform each item `i -> mapper(i)`.
    2. **Reduce**: combine all mapped results using `reducer`.
- You can implement it with a simple loop:
    - Initialize some accumulator with a sensible “zero value.”
    - For each element in inputs, apply `mapper` first, then incorporate the mapped value into your accumulator using `reducer`.
- Decide on how to handle the **empty** `inputs`. The typical approach might be to return `0` for operations like sum, or define a special case for multiplication, etc.


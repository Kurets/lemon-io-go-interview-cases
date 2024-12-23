# Pipeline with Generator and Filter

## Task Description

Implement a **pipeline** that consists of the following stages:

1. **Generator**:
    - Sends a sequence of integers into a channel.
    - Operates within its own goroutine.

2. **Filter**:
    - Reads integers from an input channel and filters them based on a condition (e.g., keep only even numbers).
    - Operates within its own goroutine and sends the filtered results into an output channel.

3. **Consumer**:
    - Reads integers from the filtered output channel.

These stages should be composed using **channels** to demonstrate a pipeline pattern in concurrent Go programming.

### Function Signatures

```go
package main

func generator(nums ...int) <-chan int {
	// Your code here
	return nil
}

func filterEven(in <-chan int) <-chan int {
	// Your code here
	return nil
}
```

### Requirements

1. **Generator**:
    - Accepts a variadic number of integers (`nums ...int`).
    - Sends each integer into a channel, then closes the channel.

2. **Filter**:
    - Accepts an input channel (`<-chan int`).
    - Sends only even numbers into an output channel, then closes the channel.

3. **Concurrency**:
    - Each stage (generator and filter) must run in its own goroutine.
    - Ensure channels are properly closed when processing is complete.

4. **Composition**:
    - Combine these stages into a pipeline where the generator feeds the filter, and the filter feeds the consumer.

### Example Usage

```go
package main

import "fmt"

func main() {
	// Create a pipeline
	gen := generator(1, 2, 3, 4, 5, 6)
	even := filterEven(gen)

	// Consume the filtered output
	for num := range even {
		fmt.Println(num) // Output: 2, 4, 6
	}

}
```

### Notes

- Use **channels** to connect the stages of the pipeline.
- Each stage should close its output channel when done.
- Ensure the pipeline can handle empty input gracefully.

### Hints

- **Generator**:
    - Create a channel and launch a goroutine to send values into it.
    - Close the channel when all values are sent.

- **Filter**:
    - Create an output channel and launch a goroutine to process values from the input channel.
    - Close the output channel after all input values are processed.

- **Channel Communication**:
    - Use `range` to iterate over values from a channel.
    - Ensure channels are closed to signal completion to downstream stages.

# Merging Sorted Channels

## Task Description

You need to implement a function:

```go
func mergeSortedChannels(chs []<-chan int) <-chan int {
    // Your code here
    return nil
}

```

## Requirements
- You are given multiple channels, each producing sorted integers in non-decreasing order.
- Your task is to merge these channels into one output channel that yields all integers from the input channels in a globally sorted order.
- Once you finish reading all input channels, you must close the output channel so that consumers can range over it until completion.

## Considerations
- Because each input channel is sorted, you can think of this as merging multiple sorted lists (but in a streaming fashion).
- You need to ensure that no integer is lost or duplicated.
- The final output should be sorted in non-decreasing order.
- Handle edge cases gracefully, including:
  - Empty channels (some or all input channels might not produce any values). 
  - Single channel (merging a single sorted channel is a trivial pass-through). 
  - Channels of differing lengths. 
  - Duplicate values across channels should appear correctly in sorted order.

## Example Usage
```go
func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)
    ch3 := make(chan int)

    // Start producers
    go func() {
        // Sorted values
        for _, v := range []int{1, 4, 7, 10} {
            ch1 <- v
        }
        close(ch1)
    }()

    go func() {
        // Sorted values
        for _, v := range []int{2, 2, 6, 9} {
            ch2 <- v
        }
        close(ch2)
    }()

    go func() {
        // Sorted values
        for _, v := range []int{3, 5} {
            ch3 <- v
        }
        close(ch3)
    }()

    // Merge
    merged := mergeSortedChannels([]<-chan int{ch1, ch2, ch3})

    // Consume merged results
    for val := range merged {
        fmt.Println(val)
    }
    // Output should be: 1,2,2,3,4,5,6,7,9,10
}
```

## Hints

- A typical approach might use a min-heap or priority queue (based on the smallest current head
  among the channels).
- Alternatively, you could “k-way merge” by reading from each channel in a loop, always taking the
  smallest available next integer.
- Ensure you handle closing the output channel once all input channels are exhausted.

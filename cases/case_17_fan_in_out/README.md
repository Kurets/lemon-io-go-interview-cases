# Fan-Out / Fan-In Task

You need to implement two functions in Go to demonstrate a **fan-out/fan-in** concurrency pattern:

1. `fanOut(in <-chan int, workers int) []<-chan int`  
2. `fanIn(channels []<-chan int) <-chan int`

## Task Requirements

1. **`fanOut`**  
   - **Signature**:  
     ```go
     func fanOut(in <-chan int, workers int) []<-chan int {
         // your code here
         return nil
     }
     ```
   - **Goal**:  
     - Launch **`workers`** goroutines that collectively process all items from the input channel `in`.  
     - Each worker should receive a subset of the input items (ensuring each item is processed exactly once by exactly one worker).  
     - Each worker outputs processed results to its own channel.  
   - **Return**: A slice of channels—each one corresponding to a single worker’s output stream.

2. **`fanIn`**  
   - **Signature**:  
     ```go
     func fanIn(channels []<-chan int) <-chan int {
         // your code here
         return nil
     }
     ```
   - **Goal**:  
     - Merge multiple channels (produced by the `fanOut` workers) into a single output channel.  
     - Once **all** worker output channels are fully read, the merged output channel should be closed.

## Important Details

- **Data Transformation**:  
  - Unless otherwise specified, assume the data **should not be changed**. If an integer `n` enters the `in` channel, it should appear **unchanged** on the final merged output channel.
  
- **Distribution Logic**:  
  - Each item from `in` must be processed exactly once.  
  - You can decide how to distribute items among workers (e.g., round-robin).
  
- **Channel Closure**:  
  - `fanOut` must properly close each worker’s output channel when processing is finished.  
  - `fanIn` must close its combined output channel once **all** worker channels are exhausted.

- **Parallelism**:  
  - Use goroutines to handle concurrency.  
  - Each worker should run in a separate goroutine.

## Example Usage

```go
func main() {
    in := make(chan int)
    go func() {
        for i := 1; i <= 5; i++ {
            in <- i
        }
        close(in)
    }()

    // Distribute items among 2 workers
    outs := fanOut(in, 2)

    // Merge worker outputs
    merged := fanIn(outs)

    // Print the results
    for val := range merged {
        fmt.Println(val)
    }
}
```

## Hints

- Consider creating a small “dispatcher” goroutine in `fanOut` that reads from `in` and distributes
  items across worker-specific input channels.
- Each worker goroutine should read from its **own** channel, process the data, and send results to
  its **own** output channel.
- In `fanIn`, you can use a `sync.WaitGroup` to track when all worker channels have finished sending
  data. When all are done, close the merged output channel.

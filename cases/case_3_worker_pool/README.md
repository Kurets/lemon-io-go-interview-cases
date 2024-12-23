# Worker Pool

## Task Description

Implement a **worker pool** that processes jobs from a channel using a fixed number of workers. Each worker should:

1. Read jobs from a `jobs` channel.
2. Process the job.
3. Send the results to a `results` channel.

### Function Signature

```go
package main

func startWorkerPool(numWorkers int, jobs <-chan int, results chan<- int) {
	// Your code here
}
```

### Requirements

1. **Input**:
    - `numWorkers`: The number of workers to spawn.
    - `jobs`: A channel containing jobs (integers) to be processed.

2. **Output**:
    - `results`: A channel where the results of the job processing are sent.

3. **Behavior**:
    - The worker pool should start exactly `numWorkers` goroutines.
    - Each worker should read from the `jobs` channel, process the job, and send the result to the `results` channel.
    - When the `jobs` channel is closed, all workers should stop processing.
    - Close the `results` channel once all jobs are processed.

### Example Usage

```go
package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Add jobs to the jobs channel
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)

	// Start the worker pool
	startWorkerPool(3, jobs, results)

	// Print results from the results channel
	for result := range results {
		fmt.Println(result)
	}
}
```

### Notes

- Each job can represent a task like squaring a number, performing a calculation, or any computational work.
- Ensure all workers terminate gracefully when there are no more jobs to process.
- The results channel should be closed by the worker pool once all results are sent.

### Hints

- Use a `sync.WaitGroup` to manage the lifecycle of worker goroutines.
- Ensure the `results` channel is only closed after all workers finish processing.
- Each worker should operate in its own goroutine to process jobs concurrently.

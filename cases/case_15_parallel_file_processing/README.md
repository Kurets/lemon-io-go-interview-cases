# Parallel File Processing

## Task Description

Implement a function that processes a list of files in parallel using a specified number of worker goroutines. Each
worker processes files concurrently by calling a dummy processing function on each file.

### Function Signature

```go
package main

func processFiles(files []string, workers int) {
	// Your code here
}
```

### Requirements

1. **Input**:
    - `files`: A slice of filenames (strings) to be processed.
    - `workers`: The number of worker goroutines to use for concurrent processing.

2. **Behavior**:
    - Process all files in the `files` slice by passing each to a predefined dummy function (`testProcessFunc`).
    - Distribute the workload across the specified number of worker goroutines.

3. **Concurrency**:
    - Use a worker pool pattern:
        1. Create a channel to pass filenames to workers.
        2. Spawn the specified number of workers.
        3. Ensure all files are processed and all workers exit gracefully.

4. **Edge Cases**:
    - If `files` is empty, the function should do nothing.
    - If `workers` is 0 or negative, default to using 1 worker.

### Example Usage

```go
package main

import "fmt"

func main() {
	files := []string{"file1", "file2", "file3"}

	// Example dummy function
	testProcessFunc := func(file string) {
		fmt.Println("Processing:", file)
	}

	// Process files using 3 workers
	processFiles(files, 3)

}
```

### Notes

- **Worker Pool Pattern**:
    - Create a fixed number of workers, each consuming tasks from a shared channel.
    - Close the channel after sending all files to signal workers to stop.

- **Thread Safety**:
    - If maintaining shared state (e.g., tracking processed files), use synchronization primitives like `sync.Mutex`.

- **Default Behavior**:
    - If `workers` is invalid (e.g., ≤0), default to a single worker.

### Edge Cases

1. Empty `files` slice (`files == nil` or `len(files) == 0`).
2. Negative or zero `workers`.
3. Very large number of files.
4. Long-running or blocking dummy processing function.

### Hints

- **Channel Usage**:
    - Use a channel to pass filenames to workers.
    - Close the channel after sending all filenames.

- **Worker Loop**:
    - Each worker should:
        1. Continuously read from the channel.
        2. Exit when the channel is closed.

- **Synchronization**:
    - Use `sync.WaitGroup` to ensure all workers finish before returning from the function.

### Example Workflow

1. Create a channel to pass filenames.
2. Spawn `workers` goroutines, each reading from the channel and processing files.
3. Send filenames to the channel.
4. Close the channel after all filenames are sent.
5. Wait for all workers to finish.

### Performance Considerations

- Ensure that the number of workers is optimal for the workload.
- Avoid deadlocks by using buffered channels or ensuring timely consumption of messages.

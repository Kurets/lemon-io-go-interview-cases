# Concurrency-Safe FIFO Queue

## Task Description

You need to implement a **FIFO queue** (`SafeQueue`) that supports concurrent access from multiple goroutines. The queue should offer two methods:

```go
type SafeQueue struct {
	// Your code here
}

func (q *SafeQueue) Enqueue(val interface{}) {
	// Your code here
}

func (q *SafeQueue) Dequeue() (interface{}, bool) {
	// Your code here
	return nil, false
}
```

### Requirements

1. **FIFO Behavior**
    - First item Enqueued should be the first item Dequeued.

2. **Concurrency Safety**
    - Multiple goroutines may call `Enqueue` or `Dequeue` at the same time.
    - Must avoid race conditions. Usually, this is done with a `sync.Mutex`, or a channel-based approach, or `sync.Cond`.
    - `Dequeue` must not panic if the queue is empty; it should return `(nil, false)` or an equivalent indicator.

3. **Return Values**
    - `Enqueue(val interface{})`: no return, just add `val` to the end.
    - `Dequeue() (interface{}, bool)`: returns `(theItem, true)` if successful, or `(nil, false)` if the queue is empty.

4. **Edge Cases**
    - **Empty queue** on dequeue => return `(nil, false)`.
    - Concurrency with large number of `Enqueue` and `Dequeue` calls.
    - Ensuring no data lost or duplicated.

### Example Usage

```go
package main

func main() {
	var q SafeQueue
	q.Enqueue("hello")
	q.Enqueue("world")

	val, ok := q.Dequeue()
	// val => "hello", ok => true
	val, ok = q.Dequeue()
	// val => "world", ok => true
	val, ok = q.Dequeue()
	// val => nil, ok => false (empty queue)
}
```

### Hints

- A simple approach is to store items in a slice:
    - Use a `sync.Mutex` to guard access to this slice.
    - `Enqueue` appends to the slice.
    - `Dequeue` removes from the front if not empty.
- Watch out for slice re-slicing or shifting overhead if you remove from the front. An alternative is to keep front/back indexes.
- Ensure `Dequeue` returns immediately when empty rather than blocking. (Unless a blocking approach is desired, but typically the specification just returns `(nil, false)`).  


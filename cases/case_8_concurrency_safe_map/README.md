# Concurrency-Safe Map

## Task Description

Implement a **concurrency-safe map** with the following methods:

1. **Get(key string) (string, bool)**:
    - Retrieves the value associated with the key.
    - Returns `false` if the key does not exist.

2. **Set(key, value string)**:
    - Stores the value associated with the key.
    - Overwrites the value if the key already exists.

3. **Delete(key string)**:
    - Removes the key-value pair from the map.
    - If the key does not exist, the method should handle gracefully.

### Function Signatures

```go
package main

type ConcurrentMap struct {
	// Your code here
}

func NewConcurrentMap() *ConcurrentMap {
	// Your code here
	return nil
}

func (m *ConcurrentMap) Get(key string) (string, bool) {
	// Your code here
	return "", false
}

func (m *ConcurrentMap) Set(key, value string) {
	// Your code here
}

func (m *ConcurrentMap) Delete(key string) {
	// Your code here
}
```

### Requirements

1. **Concurrency Safety**:
    - Use a `sync.RWMutex` to allow multiple readers or a single writer at a time.
    - Ensure that `Set` and `Delete` operations are exclusive and block other operations.

2. **Efficiency**:
    - Use `RLock` for read operations to allow multiple readers concurrently.
    - Use `Lock` for write operations to ensure thread safety.

3. **Edge Cases**:
    - Handle retrieval of non-existent keys gracefully.
    - Ensure deletion of non-existent keys does not cause errors.

### Example Usage

```go
package main

import "fmt"

func main() {
	cm := NewConcurrentMap()

	cm.Set("foo", "bar")
	val, ok := cm.Get("foo")
	if ok {
		fmt.Println("Value for 'foo':", val) // Output: Value for 'foo': bar
	}

	cm.Delete("foo")
	_, ok = cm.Get("foo")
	if !ok {
		fmt.Println("'foo' was deleted.") // Output: 'foo' was deleted.
	}

}
```

### Notes

- Use a `map[string]string` as the underlying storage for key-value pairs.
- Protect all operations on the map with a `sync.RWMutex` to ensure thread safety.
- The `Delete` method should handle keys that do not exist gracefully.

### Hints

- **Get Method**:
    - Use `RLock` to allow multiple concurrent readers.
    - Ensure you release the lock after retrieving the value.

- **Set Method**:
    - Use `Lock` to prevent other readers or writers from accessing the map while updating it.

- **Delete Method**:
    - Use `Lock` to ensure exclusive access while removing a key.

- **Initialization**:
    - Provide a constructor `NewConcurrentMap` to initialize the map and mutex properly.

# Pub/Sub System

## Task Description

Implement a simple **Publish/Subscribe (Pub/Sub)** system with the following features:

1. Subscribers can register to receive messages by subscribing.
2. Publishers can send messages, which are delivered to all subscribers.
3. The system should support **concurrent publishers** and **subscribers** safely.

### Function Signatures

```go
package main

type PubSub struct {
	// Your code here
}

func NewPubSub() *PubSub {
	// Your code here
	return nil
}

func (ps *PubSub) Subscribe() <-chan string {
	// Your code here
	return nil
}

func (ps *PubSub) Publish(msg string) {
	// Your code here
}
```

### Requirements

1. **Concurrency Safety**:
    - Ensure thread-safe access to the list of subscribers.
    - Use synchronization primitives such as `sync.RWMutex`.

2. **Subscriber Channels**:
    - Each subscriber receives messages on their own channel.
    - Use buffered channels to avoid deadlocks during publishing.

3. **Publishing**:
    - Deliver each message to all current subscribers.
    - Subscribers added after a message is published should not receive that message.

4. **Scalability**:
    - Support multiple concurrent publishers and subscribers efficiently.

### Example Usage

```go
package main

import "fmt"

func main() {
	ps := NewPubSub()

	sub1 := ps.Subscribe()
	sub2 := ps.Subscribe()

	go func() {
		ps.Publish("hello")
	}()

	fmt.Println(<-sub1) // Output: "hello"
	fmt.Println(<-sub2) // Output: "hello"

}
```

### Notes

- **Subscription**:
    - A new subscriber is added by creating a new channel and appending it to the list of subscribers.

- **Publishing**:
    - On publishing a message, iterate over the list of subscribers and send the message to each subscriber’s channel.

- **Channel Management**:
    - Ensure proper handling of channels (e.g., closing channels when necessary).

### Edge Cases

1. Publishing messages when there are no subscribers.
2. Handling subscribers that are added after messages have been published.
3. Ensuring no subscriber misses a message when multiple publishers publish concurrently.
4. Subscribers unsubscribing or being garbage collected (if implemented).

### Hints

- **Mutex Usage**:
    - Use `sync.RWMutex` for protecting the subscriber list:
        - `RLock` for read-only access (e.g., publishing).
        - `Lock` for modifying the list (e.g., subscribing).

- **Channel Buffering**:
    - Use buffered channels for subscriber channels to prevent deadlocks if a subscriber is slow to consume messages.

- **Iterating Over Subscribers**:
    - Ensure thread safety when iterating over the list of subscribers during publishing.

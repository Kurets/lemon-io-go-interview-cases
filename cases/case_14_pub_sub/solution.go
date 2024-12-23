package case_14_pub_sub

import (
	"sync"
)

type PubSub struct {
	mu          sync.RWMutex
	subscribers []chan string
}

func NewPubSub() *PubSub {
	return &PubSub{}
}

func (ps *PubSub) Subscribe() <-chan string {
	ch := make(chan string, 10)
	ps.mu.Lock()
	ps.subscribers = append(ps.subscribers, ch)
	ps.mu.Unlock()
	return ch
}

func (ps *PubSub) Publish(msg string) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	for _, ch := range ps.subscribers {
		ch <- msg
	}
}

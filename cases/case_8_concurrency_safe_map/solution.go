package case_8_concurrency_safe_map

import (
	"fmt"
	"sync"
)

type ConcurrentMap struct {
	mu sync.RWMutex
	m  map[string]string
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		m: make(map[string]string),
	}
}

func (m *ConcurrentMap) Get(key string) (string, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	val, ok := m.m[key]
	return val, ok
}

func (m *ConcurrentMap) Set(key, value string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m[key] = value
}

func (m *ConcurrentMap) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.m[key]; ok {
		delete(m.m, key)
		fmt.Printf("Deleted key: %s\n", key)
	}
}

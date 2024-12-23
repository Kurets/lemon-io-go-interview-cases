package case_2_lru_cache

import (
	"container/list"
	"sync"
)

type entry struct {
	key, value string
}

type LRUCache struct {
	capacity int
	mu       sync.Mutex
	ll       *list.List
	cache    map[string]*list.Element
}

func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		capacity: cap,
		ll:       list.New(),
		cache:    make(map[string]*list.Element),
	}
}

func (c *LRUCache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if elem, ok := c.cache[key]; ok {
		c.ll.MoveToFront(elem)
		return elem.Value.(*entry).value, true
	}
	return "", false
}

func (c *LRUCache) Put(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if elem, ok := c.cache[key]; ok {
		c.ll.MoveToFront(elem)
		elem.Value.(*entry).value = value
		return
	}

	newEntry := &entry{key: key, value: value}
	elem := c.ll.PushFront(newEntry)
	c.cache[key] = elem

	if c.ll.Len() > c.capacity {
		last := c.ll.Back()
		c.ll.Remove(last)
		delete(c.cache, last.Value.(*entry).key)
	}
}

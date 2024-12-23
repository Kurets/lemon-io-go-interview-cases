package case_5_memoization

import (
	"sync"
)

func Memoize(f func(string) string) func(string) string {
	var mu sync.Mutex
	cache := make(map[string]string)
	return func(key string) string {
		mu.Lock()
		if val, ok := cache[key]; ok {
			mu.Unlock()
			return val
		}
		mu.Unlock()
		val := f(key)
		mu.Lock()
		cache[key] = val
		mu.Unlock()
		return val
	}
}

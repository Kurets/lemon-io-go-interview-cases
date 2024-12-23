package case_8_concurrency_safe_map

import (
	"fmt"
	"sync"
	"testing"
)

func TestBasicOperations(t *testing.T) {
	cm := NewConcurrentMap()

	cm.Set("foo", "bar")
	val, ok := cm.Get("foo")
	if !ok || val != "bar" {
		t.Errorf("Expected Get('foo') = 'bar', true; got '%s', %v", val, ok)
	}

	cm.Delete("foo")
	val, ok = cm.Get("foo")
	if ok {
		t.Errorf("Expected Get('foo') to return false after deletion; got '%s', %v", val, ok)
	}
}

func TestNonExistentKey(t *testing.T) {
	cm := NewConcurrentMap()

	val, ok := cm.Get("nonexistent")
	if ok {
		t.Errorf("Expected Get('nonexistent') to return false, got true with value '%s'", val)
	}

	cm.Delete("nonexistent")
}

func TestMultipleReadersSingleWriter(t *testing.T) {
	cm := NewConcurrentMap()
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 50; i++ {
			cm.Set("key", string(rune('A'+i)))
		}
	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 50; j++ {
				cm.Get("key")
			}
		}()
	}

	wg.Wait()
}

func TestDeleteConsistency(t *testing.T) {
	cm := NewConcurrentMap()

	keys := []string{"A", "B", "C", "D"}
	for _, key := range keys {
		cm.Set(key, "value")
	}

	for _, key := range keys {
		cm.Delete(key)
	}

	for _, key := range keys {
		if _, ok := cm.Get(key); ok {
			t.Errorf("Expected key '%s' to be deleted, but it was found in map", key)
		}
	}
}

func TestConcurrentAccess(t *testing.T) {
	cm := NewConcurrentMap()
	var wg sync.WaitGroup

	numKeys := 100
	keys := make([]string, numKeys)
	for i := 0; i < numKeys; i++ {
		keys[i] = fmt.Sprintf("key-%d", i)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, key := range keys {
			cm.Set(key, "value")
		}
	}()
	wg.Wait()

	wg.Add(2)
	go func() {
		defer wg.Done()
		for _, key := range keys {
			_, _ = cm.Get(key)
		}
	}()
	go func() {
		defer wg.Done()
		for _, key := range keys {
			cm.Delete(key)
		}
	}()
	wg.Wait()

	for _, key := range keys {
		if _, ok := cm.Get(key); ok {
			t.Errorf("Expected key '%s' to be deleted, but found in map", key)
		}
	}
}

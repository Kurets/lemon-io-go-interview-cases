package case_5_memoization

import (
	"sync"
	"testing"
)

func TestCachingBehavior(t *testing.T) {
	callCount := 0
	testFunc := func(s string) string {
		callCount++
		return s + "-computed"
	}

	memoized := Memoize(testFunc)

	result := memoized("test")
	if result != "test-computed" {
		t.Errorf("Expected 'test-computed', got %s", result)
	}

	result = memoized("test")
	if result != "test-computed" {
		t.Errorf("Expected 'test-computed' on cached call, got %s", result)
	}

	if callCount != 1 {
		t.Errorf("Expected testFunc to be called once, but was called %d times", callCount)
	}
}

func TestCorrectnessWithMultipleKeys(t *testing.T) {
	callCount := 0
	testFunc := func(s string) string {
		callCount++
		return s + "-computed"
	}

	memoized := Memoize(testFunc)

	result1 := memoized("key1")
	result2 := memoized("key2")
	if result1 != "key1-computed" || result2 != "key2-computed" {
		t.Errorf("Expected 'key1-computed' and 'key2-computed', got %s and %s", result1, result2)
	}

	if callCount != 2 {
		t.Errorf("Expected testFunc to be called twice, but was called %d times", callCount)
	}

	result1 = memoized("key1")
	result2 = memoized("key2")
	if callCount != 2 {
		t.Errorf("Expected cached calls not to increase call count, but got %d calls", callCount)
	}
}

func TestConcurrentAccess(t *testing.T) {
	callCount := 0
	testFunc := func(s string) string {
		callCount++
		return s + "-computed"
	}

	memoized := Memoize(testFunc)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if memoized("test") != "test-computed" {
				t.Errorf("Expected 'test-computed', got something else")
			}
		}()
	}

	wg.Wait()

	if callCount != 1 {
		t.Errorf("Expected testFunc to be called once, but was called %d times", callCount)
	}
}

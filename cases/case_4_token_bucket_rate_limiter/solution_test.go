package case_4_token_bucket_rate_limiter

import (
	"sync"
	"testing"
	"time"
)

func TestAllowInitialTokens(t *testing.T) {
	rl := NewRateLimiter(3)
	defer close(rl.stopChan)

	for i := 0; i < 3; i++ {
		if !rl.Allow() {
			t.Errorf("Expected Allow() to return true, but got false on iteration %d", i)
		}
	}

	if rl.Allow() {
		t.Errorf("Expected Allow() to return false after consuming all tokens, but got true")
	}
}

func TestRefillBehavior(t *testing.T) {
	rl := NewRateLimiter(2)
	defer close(rl.stopChan)

	if !rl.Allow() || !rl.Allow() {
		t.Errorf("Expected initial Allow() calls to return true")
	}

	if rl.Allow() {
		t.Errorf("Expected Allow() to return false after exhausting tokens")
	}

	time.Sleep(1100 * time.Millisecond)

	if !rl.Allow() {
		t.Errorf("Expected Allow() to return true after tokens are refilled")
	}
}

func TestTokenExhaustionAndRefill(t *testing.T) {
	rl := NewRateLimiter(1)
	defer close(rl.stopChan)

	if !rl.Allow() {
		t.Errorf("Expected Allow() to return true for the first call")
	}

	if rl.Allow() {
		t.Errorf("Expected Allow() to return false after consuming the only token")
	}

	time.Sleep(1100 * time.Millisecond)

	if !rl.Allow() {
		t.Errorf("Expected Allow() to return true after refill")
	}
}

func TestConcurrentAccess(t *testing.T) {
	rl := NewRateLimiter(5)
	defer close(rl.stopChan)

	var wg sync.WaitGroup
	successCount := 0
	mutex := sync.Mutex{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if rl.Allow() {
				mutex.Lock()
				successCount++
				mutex.Unlock()
			}
		}()
	}

	wg.Wait()

	if successCount > 5 {
		t.Errorf("Expected at most 5 successful Allow() calls, but got %d", successCount)
	}
}

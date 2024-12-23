package case_4_token_bucket_rate_limiter

import (
	"sync"
	"time"
)

type RateLimiter struct {
	rate     int
	tokens   int
	mu       sync.Mutex
	ticker   *time.Ticker
	stopChan chan struct{}
}

func NewRateLimiter(rate int) *RateLimiter {
	rl := &RateLimiter{
		rate:     rate,
		tokens:   rate,
		ticker:   time.NewTicker(time.Second),
		stopChan: make(chan struct{}),
	}
	go rl.refiller()
	return rl
}

func (r *RateLimiter) refiller() {
	for {
		select {
		case <-r.ticker.C:
			r.mu.Lock()
			r.tokens = r.rate
			r.mu.Unlock()
		case <-r.stopChan:
			return
		}
	}
}

func (r *RateLimiter) Allow() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.tokens > 0 {
		r.tokens--
		return true
	}
	return false
}

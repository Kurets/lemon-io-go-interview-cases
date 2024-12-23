package case_14_pub_sub

import (
	"sync"
	"testing"
	"time"
)

func TestSingleSubscriberSingleMessage(t *testing.T) {
	ps := NewPubSub()
	sub := ps.Subscribe()

	ps.Publish("hello")

	select {
	case msg := <-sub:
		if msg != "hello" {
			t.Errorf("expected 'hello', got '%s'", msg)
		}
	case <-time.After(1 * time.Second):
		t.Error("timed out waiting for message")
	}
}

func TestMultipleSubscribers(t *testing.T) {
	ps := NewPubSub()
	sub1 := ps.Subscribe()
	sub2 := ps.Subscribe()

	ps.Publish("world")

	for i, sub := range []<-chan string{sub1, sub2} {
		select {
		case msg := <-sub:
			if msg != "world" {
				t.Errorf("subscriber %d expected 'world', got '%s'", i, msg)
			}
		case <-time.After(1 * time.Second):
			t.Errorf("subscriber %d timed out waiting for message", i)
		}
	}
}

func TestSubscribeAfterPublish(t *testing.T) {
	ps := NewPubSub()
	ps.Publish("before")

	sub := ps.Subscribe()

	ps.Publish("after")

	select {
	case msg := <-sub:
		if msg != "after" {
			t.Errorf("expected 'after', got '%s'", msg)
		}
	case <-time.After(1 * time.Second):
		t.Error("timed out waiting for message")
	}

	select {
	case msg := <-sub:
		t.Errorf("unexpected message '%s' received, should not have gotten 'before'", msg)
	default:
	}
}

func TestConcurrentPublishing(t *testing.T) {
	ps := NewPubSub()
	sub := ps.Subscribe()

	messages := []string{"msg1", "msg2", "msg3", "msg4", "msg5"}

	var wg sync.WaitGroup
	for _, m := range messages {
		wg.Add(1)
		go func(msg string) {
			defer wg.Done()
			ps.Publish(msg)
		}(m)
	}

	wg.Wait()

	received := make(map[string]bool)
	for i := 0; i < len(messages); i++ {
		select {
		case msg := <-sub:
			received[msg] = true
		case <-time.After(1 * time.Second):
			t.Error("timed out waiting for messages")
			return
		}
	}

	for _, m := range messages {
		if !received[m] {
			t.Errorf("subscriber did not receive message '%s'", m)
		}
	}
}

func TestNoSubscribers(t *testing.T) {
	ps := NewPubSub()

	ps.Publish("no_subs_message")

	sub := ps.Subscribe()
	ps.Publish("new_message")

	select {
	case msg := <-sub:
		if msg != "new_message" {
			t.Errorf("expected 'new_message', got '%s'", msg)
		}
	case <-time.After(1 * time.Second):
		t.Error("timed out waiting for message")
	}
}

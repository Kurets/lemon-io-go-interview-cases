package case_30_round_robin

import (
	"testing"
)

func TestRoundRobinBasic(t *testing.T) {
	endpoints := []string{"server1", "server2", "server3"}
	rr := NewRoundRobin(endpoints)

	if got := rr.Next(); got != "server1" {
		t.Errorf("expected 'server1', got '%s'", got)
	}

	if got := rr.Next(); got != "server2" {
		t.Errorf("expected 'server2', got '%s'", got)
	}

	if got := rr.Next(); got != "server3" {
		t.Errorf("expected 'server3', got '%s'", got)
	}

	if got := rr.Next(); got != "server1" {
		t.Errorf("expected 'server1' again, got '%s'", got)
	}
}

func TestRoundRobinCycle(t *testing.T) {
	endpoints := []string{"s1", "s2"}
	rr := NewRoundRobin(endpoints)

	var sequence []string
	for i := 0; i < 6; i++ {
		sequence = append(sequence, rr.Next())
	}

	expected := []string{"s1", "s2", "s1", "s2", "s1", "s2"}
	for i := range sequence {
		if sequence[i] != expected[i] {
			t.Errorf("call #%d => expected '%s', got '%s'", i+1, expected[i], sequence[i])
		}
	}
}

func TestRoundRobinSingleEndpoint(t *testing.T) {
	endpoints := []string{"onlyServer"}
	rr := NewRoundRobin(endpoints)

	for i := 0; i < 5; i++ {
		got := rr.Next()
		if got != "onlyServer" {
			t.Errorf("expected 'onlyServer', got '%s'", got)
		}
	}
}

func TestRoundRobinEmpty(t *testing.T) {
	var endpoints []string
	rr := NewRoundRobin(endpoints)

	for i := 0; i < 3; i++ {
		got := rr.Next()
		if got != "" {
			t.Errorf("expected empty string for empty endpoints, got '%s'", got)
		}
	}
}

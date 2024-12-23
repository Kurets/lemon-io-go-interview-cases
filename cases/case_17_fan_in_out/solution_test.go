package case_17_fan_in_out

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestFanOutFanInBasic(t *testing.T) {
	in := make(chan int)
	workers := 3

	outs := fanOut(in, workers)
	merged := fanIn(outs)

	inputs := []int{1, 2, 3, 4, 5, 10, 20}
	go func() {
		for _, v := range inputs {
			in <- v
		}
		close(in)
	}()

	var results []int
	done := make(chan struct{})

	go func() {
		for val := range merged {
			results = append(results, val)
		}
		close(done)
	}()

	select {
	case <-done:
		sort.Ints(inputs)
		sort.Ints(results)

		if !reflect.DeepEqual(inputs, results) {
			t.Errorf("expected %v, got %v", inputs, results)
		}
	case <-time.After(2 * time.Second):
		t.Error("test timed out, possibly fanIn not closing or not all items processed")
	}
}

func TestFanOutFanInEmpty(t *testing.T) {
	in := make(chan int)
	outs := fanOut(in, 4)
	merged := fanIn(outs)

	close(in)

	var results []int
	done := make(chan struct{})
	go func() {
		for val := range merged {
			results = append(results, val)
		}
		close(done)
	}()

	select {
	case <-done:
		if len(results) != 0 {
			t.Errorf("expected no results, got %v", results)
		}
	case <-time.After(time.Second):
		t.Error("test timed out waiting for empty output")
	}
}

func TestFanOutFanInSingle(t *testing.T) {
	in := make(chan int)
	outs := fanOut(in, 2)
	merged := fanIn(outs)

	go func() {
		in <- 42
		close(in)
	}()

	var results []int
	done := make(chan struct{})
	go func() {
		for val := range merged {
			results = append(results, val)
		}
		close(done)
	}()

	select {
	case <-done:
		if len(results) != 1 || results[0] != 42 {
			t.Errorf("expected [42], got %v", results)
		}
	case <-time.After(time.Second):
		t.Error("test timed out")
	}
}

func TestFanOutFanInMultipleWorkers(t *testing.T) {
	in := make(chan int)
	workers := 5
	count := 100
	outs := fanOut(in, workers)
	merged := fanIn(outs)

	go func() {
		for i := 0; i < count; i++ {
			in <- i
		}
		close(in)
	}()

	resultsMap := make(map[int]bool)
	done := make(chan struct{})
	go func() {
		for val := range merged {
			resultsMap[val] = true
		}
		close(done)
	}()

	select {
	case <-done:
		if len(resultsMap) != count {
			t.Errorf("expected %d unique items, got %d", count, len(resultsMap))
		}
		for i := 0; i < count; i++ {
			if !resultsMap[i] {
				t.Errorf("missing value %d in output", i)
			}
		}
	case <-time.After(3 * time.Second):
		t.Error("test timed out waiting for all results")
	}
}

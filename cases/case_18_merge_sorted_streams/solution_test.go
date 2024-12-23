package case_18_merge_sorted_streams

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

func collectChannel(ch <-chan int) []int {
	var result []int
	for v := range ch {
		result = append(result, v)
	}
	return result
}

func TestSingleChannel(t *testing.T) {
	ch := make(chan int)
	go func() {
		for _, v := range []int{1, 2, 3} {
			ch <- v
		}
		close(ch)
	}()

	out := mergeSortedChannels([]<-chan int{ch})
	result := collectChannel(out)

	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("single channel test failed: expected %v, got %v", expected, result)
	}
}

func TestMultipleChannels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		for _, v := range []int{1, 4, 7, 10} {
			ch1 <- v
		}
		close(ch1)
	}()

	go func() {
		for _, v := range []int{2, 2, 6, 9} {
			ch2 <- v
		}
		close(ch2)
	}()

	go func() {
		for _, v := range []int{3, 5} {
			ch3 <- v
		}
		close(ch3)
	}()

	out := mergeSortedChannels([]<-chan int{ch1, ch2, ch3})
	result := collectChannel(out)

	expected := []int{1, 2, 2, 3, 4, 5, 6, 7, 9, 10}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("multi-channel test failed: expected %v, got %v", expected, result)
	}
}

func TestDifferentLengths(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for _, v := range []int{1, 10} {
			ch1 <- v
		}
		close(ch1)
	}()

	go func() {
		for _, v := range []int{2, 3, 4, 5, 6, 7} {
			ch2 <- v
		}
		close(ch2)
	}()

	out := mergeSortedChannels([]<-chan int{ch1, ch2})
	result := collectChannel(out)

	expected := []int{1, 2, 3, 4, 5, 6, 7, 10}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("diff length test failed: expected %v, got %v", expected, result)
	}
}

func TestEmptyChannels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		// no data
		close(ch1)
	}()
	go func() {
		for _, v := range []int{1, 2, 3} {
			ch2 <- v
		}
		close(ch2)
	}()

	out := mergeSortedChannels([]<-chan int{ch1, ch2})
	result := collectChannel(out)

	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("empty channel test failed: expected %v, got %v", expected, result)
	}
}

func TestDuplicatesAcrossChannels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for _, v := range []int{1, 2, 2, 5} {
			ch1 <- v
		}
		close(ch1)
	}()

	go func() {
		for _, v := range []int{2, 2, 2, 6} {
			ch2 <- v
		}
		close(ch2)
	}()

	out := mergeSortedChannels([]<-chan int{ch1, ch2})
	result := collectChannel(out)

	expected := []int{1, 2, 2, 2, 2, 2, 5, 6}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("duplicate test failed: expected %v, got %v", expected, result)
	}
}

func TestAllChannelsEmpty(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() { close(ch1) }()
	go func() { close(ch2) }()
	go func() { close(ch3) }()

	out := mergeSortedChannels([]<-chan int{ch1, ch2, ch3})
	result := collectChannel(out)

	if len(result) != 0 {
		t.Errorf("all empty channels test failed: expected [], got %v", result)
	}
}

func TestTimeout(t *testing.T) {
	ch1 := make(chan int)
	go func() {
		for _, v := range []int{1, 2, 3} {
			ch1 <- v
		}
		close(ch1)
	}()

	out := mergeSortedChannels([]<-chan int{ch1})

	done := make(chan struct{})
	var result []int
	go func() {
		result = collectChannel(out)
		close(done)
	}()

	select {
	case <-done:
		expected := []int{1, 2, 3}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("timeout test failed: expected %v, got %v", expected, result)
		}
	case <-time.After(time.Second):
		t.Error("test timed out; function might be stuck")
	}
}

func TestGlobalSorting(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer close(ch1)
		for _, v := range []int{1, 3, 3, 8} {
			ch1 <- v
		}
	}()
	go func() {
		defer close(ch2)
		for _, v := range []int{2, 2, 9} {
			ch2 <- v
		}
	}()
	go func() {
		defer close(ch3)
		for _, v := range []int{4, 5, 5, 6} {
			ch3 <- v
		}
	}()

	out := mergeSortedChannels([]<-chan int{ch1, ch2, ch3})
	result := collectChannel(out)

	if !sort.IntsAreSorted(result) {
		t.Errorf("global sorting test failed: output is not sorted: %v", result)
	}
}

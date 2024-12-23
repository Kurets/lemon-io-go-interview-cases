package case_28_fenwick_tree

import (
	"testing"
)

func TestFenwickBasic(t *testing.T) {
	f := NewFenwick(5)

	for i := 1; i <= 5; i++ {
		if got := f.PrefixSum(i); got != 0 {
			t.Errorf("initial PrefixSum(%d) = %d, want 0", i, got)
		}
	}

	f.Update(1, 2)
	if got := f.PrefixSum(1); got != 2 {
		t.Errorf("PrefixSum(1) = %d, want 2", got)
	}
	for i := 2; i <= 5; i++ {
		if got := f.PrefixSum(i); got != 2 {
			t.Errorf("PrefixSum(%d) = %d, want 2", i, got)
		}
	}

	f.Update(2, 3)
	if got := f.PrefixSum(1); got != 2 {
		t.Errorf("PrefixSum(1) = %d, want 2", got)
	}
	if got := f.PrefixSum(2); got != 5 {
		t.Errorf("PrefixSum(2) = %d, want 5", got)
	}
	for i := 3; i <= 5; i++ {
		if got := f.PrefixSum(i); got != 5 {
			t.Errorf("PrefixSum(%d) = %d, want 5", i, got)
		}
	}
}

func TestFenwickLargerUpdates(t *testing.T) {
	f := NewFenwick(6)

	f.Update(1, 1)
	f.Update(2, 2)
	f.Update(3, 3)
	f.Update(5, 5)

	expected := []int{1, 3, 6, 6, 11, 11}
	for i := 1; i <= 6; i++ {
		got := f.PrefixSum(i)
		if got != expected[i-1] {
			t.Errorf("PrefixSum(%d) = %d, want %d", i, got, expected[i-1])
		}
	}
}

func TestFenwickMultipleUpdatesSameIndex(t *testing.T) {
	f := NewFenwick(5)

	f.Update(3, 2)
	f.Update(3, 4)

	if got := f.PrefixSum(2); got != 0 {
		t.Errorf("PrefixSum(2) after updates = %d, want 0", got)
	}
	for i := 3; i <= 5; i++ {
		got := f.PrefixSum(i)
		if got != 6 {
			t.Errorf("PrefixSum(%d) = %d, want 6", i, got)
		}
	}
}

func TestFenwickEdgeCases(t *testing.T) {
	f1 := NewFenwick(1)
	f1.Update(1, 10)
	if got := f1.PrefixSum(1); got != 10 {
		t.Errorf("PrefixSum(1) = %d, want 10", got)
	}

	f2 := NewFenwick(3)
	f2.Update(2, -5)
	if got := f2.PrefixSum(1); got != 0 {
		t.Errorf("PrefixSum(1) = %d, want 0 for negative update test", got)
	}
	if got := f2.PrefixSum(2); got != -5 {
		t.Errorf("PrefixSum(2) = %d, want -5", got)
	}
	if got := f2.PrefixSum(3); got != -5 {
		t.Errorf("PrefixSum(3) = %d, want -5", got)
	}
}

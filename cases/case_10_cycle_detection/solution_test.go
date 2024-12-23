package case_10_cycle_detection

import "testing"

func TestEmptyList(t *testing.T) {
	if hasCycle(nil) {
		t.Errorf("Expected hasCycle(nil) = false, got true")
	}
}

func TestSingleNodeNoCycle(t *testing.T) {
	node := &ListNode{Val: 1}
	if hasCycle(node) {
		t.Errorf("Expected hasCycle(single node with no cycle) = false, got true")
	}
}

func TestSingleNodeWithCycle(t *testing.T) {
	node := &ListNode{Val: 1}
	node.Next = node
	if !hasCycle(node) {
		t.Errorf("Expected hasCycle(single node with cycle) = true, got false")
	}
}

func TestMultipleNodesNoCycle(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n1.Next = n2
	n2.Next = n3

	if hasCycle(n1) {
		t.Errorf("Expected hasCycle(multiple nodes without cycle) = false, got true")
	}
}

func TestMultipleNodesWithCycle(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n1

	if !hasCycle(n1) {
		t.Errorf("Expected hasCycle(multiple nodes with cycle) = true, got false")
	}
}

func TestCycleAtBeginning(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n1

	if !hasCycle(n1) {
		t.Errorf("Expected hasCycle(cycle at beginning) = true, got false")
	}
}

func TestCycleInMiddle(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n2

	if !hasCycle(n1) {
		t.Errorf("Expected hasCycle(cycle in middle) = true, got false")
	}
}

package case_6_binary_search_tree

import "testing"

func TestBasicInsertAndSearch(t *testing.T) {
	var tree BST

	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)

	if !tree.Search(5) {
		t.Errorf("Expected Search(5) to return true, but got false")
	}
	if !tree.Search(3) {
		t.Errorf("Expected Search(3) to return true, but got false")
	}
	if !tree.Search(7) {
		t.Errorf("Expected Search(7) to return true, but got false")
	}

	if tree.Search(10) {
		t.Errorf("Expected Search(10) to return false, but got true")
	}
	if tree.Search(1) {
		t.Errorf("Expected Search(1) to return false, but got true")
	}
}

func TestDuplicateInsertion(t *testing.T) {
	var tree BST

	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(3)

	if tree.left != nil && tree.left.val == 3 && tree.left.left != nil {
		t.Errorf("Duplicate value 3 should not create a new node")
	}
}

func TestEmptyTreeSearch(t *testing.T) {
	var tree *BST

	if tree.Search(5) {
		t.Errorf("Expected Search(5) on nil tree to return false, but got true")
	}
}

func TestStructureIntegrity(t *testing.T) {
	var tree BST

	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(4)
	tree.Insert(6)

	if tree.val != 5 {
		t.Errorf("Root value should be 5, got %d", tree.val)
	}
	if tree.left == nil || tree.left.val != 3 {
		t.Errorf("Left child of root should be 3, got %v", tree.left)
	}
	if tree.right == nil || tree.right.val != 7 {
		t.Errorf("Right child of root should be 7, got %v", tree.right)
	}
	if tree.left.right == nil || tree.left.right.val != 4 {
		t.Errorf("Right child of left child should be 4, got %v", tree.left.right)
	}
	if tree.right.left == nil || tree.right.left.val != 6 {
		t.Errorf("Left child of right child should be 6, got %v", tree.right.left)
	}
}

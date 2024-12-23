package case_6_binary_search_tree

type BST struct {
	val   int
	left  *BST
	right *BST
}

func (b *BST) Insert(val int) {
	if b == nil {
		return
	}
	if b.val == 0 && b.left == nil && b.right == nil {
		b.val = val
		return
	}
	if val < b.val {
		if b.left == nil {
			b.left = &BST{val: val}
		} else {
			b.left.Insert(val)
		}
	} else if val > b.val {
		if b.right == nil {
			b.right = &BST{val: val}
		} else {
			b.right.Insert(val)
		}
	}
}

func (b *BST) Search(val int) bool {
	if b == nil {
		return false
	}
	if val == b.val {
		return true
	}
	if val < b.val {
		return b.left.Search(val)
	}
	return b.right.Search(val)
}

package case_27_trie

type TrieNode struct {
	children map[rune]*TrieNode
	end      bool
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) Insert(word string) {
	if t.root == nil {
		t.root = &TrieNode{children: make(map[rune]*TrieNode)}
	}
	cur := t.root
	for _, ch := range word {
		if cur.children[ch] == nil {
			cur.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		cur = cur.children[ch]
	}
	cur.end = true
}

func (t *Trie) Search(word string) bool {
	if t.root == nil {
		return false
	}
	cur := t.root
	for _, ch := range word {
		if cur.children[ch] == nil {
			return false
		}
		cur = cur.children[ch]
	}
	return cur.end
}

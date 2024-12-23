package case_27_trie

import "testing"

func TestTrieBasicInsertSearch(t *testing.T) {
	trie := &Trie{}

	trie.Insert("cat")
	found := trie.Search("cat")
	if !found {
		t.Errorf("Search('cat') = false; want true")
	}

	notFound := trie.Search("car")
	if notFound {
		t.Errorf("Search('car') = true; want false")
	}
}

func TestTrieMultipleInsertSearch(t *testing.T) {
	trie := &Trie{}

	words := []string{"hello", "helium", "help", "heap"}
	for _, w := range words {
		trie.Insert(w)
	}

	tests := []struct {
		query string
		want  bool
	}{
		{"hello", true},
		{"helium", true},
		{"help", true},
		{"heap", true},
		{"hel", false}, // prefix only
		{"heater", false},
		{"he", false},
		{"world", false},
	}

	for _, tt := range tests {
		got := trie.Search(tt.query)
		if got != tt.want {
			t.Errorf("Search('%s') = %v; want %v", tt.query, got, tt.want)
		}
	}
}

func TestTrieEmptyString(t *testing.T) {
	trie := &Trie{}
	trie.Insert("")
	if !trie.Search("") {
		t.Errorf("Empty string not found after insert")
	}
}

func TestTrieDuplicateInsert(t *testing.T) {
	trie := &Trie{}

	trie.Insert("apple")
	trie.Insert("apple")
	trie.Insert("apple")

	if !trie.Search("apple") {
		t.Errorf("Search('apple') = false; want true after multiple inserts")
	}

	if trie.Search("app") {
		t.Errorf("Search('app') = true; want false (partial match only)")
	}
}

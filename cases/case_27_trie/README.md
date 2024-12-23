# Trie Implementation

## Task Description

You need to implement a **Trie** data structure with the following methods:

```go
package main

type Trie struct {
	// Your code here
}

func (t *Trie) Insert(word string) {
	// Your code here
}

func (t *Trie) Search(word string) bool {
	// Your code here
	return false
}
```

### Requirements

1. **`Insert(word string)`**
    - Inserts the string `word` into the trie.

2. **`Search(word string) bool`**
    - Returns `true` if `word` exists in the trie, `false` otherwise.

3. **Data Structure**
    - The trie typically has a root node that may contain up to 26 children (for a basic a-z scenario) or a map of runes if handling Unicode or arbitrary characters.
    - Each child node may represent the next character in the path.
    - One common approach is to store a boolean `endOfWord` (or `isTerminal`) flag in nodes to mark complete words.

4. **Edge Cases**
    - **Empty String**: Decide if you allow an empty string to be inserted and searched. If so, handle carefully in your data structure.
    - **Duplicate Inserts**: Inserting the same word multiple times shouldn’t break anything. Searching should still return `true`.
    - **Partial Matches**: Searching for a prefix (e.g., “app” when only “apple” is inserted) should return `false` unless you consider prefixes as valid words in your design.

5. **Behavior**
    - The tests assume standard trie behavior: partial prefixes do not count as found words unless explicitly inserted.
    - `Search` must traverse the trie nodes based on each character; if any link is missing, or if the final node isn't marked as end of a word, return `false`.

### Example Usage

```go
package main

import "fmt"

func main() {
	trie := &Trie{}
	trie.Insert("hello")
	fmt.Println(trie.Search("hello")) // true
	fmt.Println(trie.Search("hel"))   // false
}
```

### Hints

- **Node Structure**:
    - A node might have a map or fixed array of children.
    - A boolean `endOfWord` to mark if a node completes a word.

- **Insert**:
    - Traverse from the root node, create child nodes if missing, then mark endOfWord on the final node.

- **Search**:
    - Traverse from the root node, following each character’s child link.
    - If at any point the link doesn’t exist, return `false`.
    - At the end, check if the final node’s `endOfWord` is `true`.


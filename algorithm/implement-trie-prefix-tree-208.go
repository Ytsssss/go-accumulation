package algorithm

type Trie struct {
	IsWord bool
	Nodes  [26]*Trie
}

func ConstructorTrie() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for i := 0; i < len(word); i++ {
		ch := word[i] - 'a'
		if node.Nodes[ch] == nil {
			node.Nodes[ch] = &Trie{}
		}
		node = node.Nodes[ch]
	}
	node.IsWord = true
}

func (this *Trie) Search(word string) bool {
	node := this
	for i := 0; i < len(word); i++ {
		ch := word[i] - 'a'
		if node.Nodes[ch] == nil {
			return false
		}
		node = node.Nodes[ch]
	}
	return node.IsWord
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for i := 0; i < len(prefix); i++ {
		ch := prefix[i] - 'a'
		if node.Nodes[ch] == nil {
			return false
		}
		node = node.Nodes[ch]
	}
	return true
}

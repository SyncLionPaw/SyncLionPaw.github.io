package main

type Trie struct {
	son [26]*Trie
	end int
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	n := this
	for _, r := range word {
		b := byte(r)
		i := b - 'a'
		if n.son[i] == nil {
			n.son[i] = &Trie{}
		}
		n = n.son[i]
	}
	n.end++
}

func (this *Trie) Search(word string) bool {
	n := this
	for _, r := range word {
		b := byte(r)
		i := b - 'a'
		if n.son[i] == nil {
			return false
		}
		n = n.son[i]
	}
	return n.end > 0 // 这里很重要，易错
}

func (this *Trie) StartsWith(prefix string) bool {
	n := this
	for _, r := range prefix {
		b := byte(r)
		i := b - 'a'
		if n.son[i] == nil {
			return false
		}
		n = n.son[i]
	}
	return true
}

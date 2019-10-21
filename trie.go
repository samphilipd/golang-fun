package main

/*
211. Add and Search Word - Data structure design
     Medium

Design a data structure that supports the following two operations:

void addWord(word) bool search(word)

search(word) can search a literal word or a regular expression string
containing only letters a-z or .. A . means it can represent any one
letter.

Example:

addWord("bad")
addWord("dad")
addWord("mad")
search("pad") -> false
search("bad") -> true
search(".ad") -> true
search("b..") -> true

Note:
You may assume that all words are consist of lowercase letters a-z.
*/
/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

type WordDictionary struct {
	Root *TrieNode
}
type TrieNode struct {
	IsWord   bool
	Children [26]*TrieNode
}

// Constructor initialize your data structure here.
func Constructor() WordDictionary {
	root := makeNode()
	root.IsWord = false
	return WordDictionary{Root: root}
}

func makeNode() *TrieNode {
	return &TrieNode{}
}

func (dict *WordDictionary) AddWord(word string) {
	addWord(dict.Root, word)
}

func addWord(root *TrieNode, word string) {
	prev := root

	for i := 0; i < len(word); i++ {
		childIdx := letterToIndex(word[i])
		curr := prev.Children[childIdx]
		if curr == nil {
			curr = makeNode()
			prev.Children[childIdx] = curr
		}

		isWord := false
		if i == len(word)-1 {
			isWord = true
		}
		curr.IsWord = isWord
		prev = curr
	}
}

// You may assume that all words are consist of lowercase letters a-z
func letterToIndex(c byte) int {
	return int(c) - 'a'
}

func (dict *WordDictionary) Search(word string) bool {
	return search(dict.Root, word)
}

func search(root *TrieNode, word string) bool {
	if root == nil {
		return false
	}
	curr := root

	for i := 0; i < len(word); i++ {
		ch := word[i]
		if ch == '.' {
			// DFS
			for l := 0; l < 26; l++ {
				if search(curr.Children[l], word[i+1:]) {
					return true
				}
			}
			return false
		} else {
			next := curr.Children[letterToIndex(ch)]
			if next != nil {
				curr = next
			} else {
				return false
			}
		}
	}
	return curr.IsWord
}

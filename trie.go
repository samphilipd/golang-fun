package main

import (
	"fmt"
	"unicode/utf8"
)

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
	Children []*LetterTrieNode
}
type LetterTrieNode struct {
	Letter   rune
	Children []*LetterTrieNode
}

// Constructor initialize your data structure here.
func Constructor() WordDictionary {
	return WordDictionary{Children: make([]*LetterTrieNode, 0)}
}

func (dict *WordDictionary) AddWord(word string) {
	if len(word) == 0 {
		return
	}
	c, _ := utf8.DecodeRuneInString(word)
	node := findNode(dict.Children, c)
	if node == nil {
		node = makeNode(c)
		dict.Children = append(dict.Children, node)
	}
	node.AddWord(word[1:])
}

// AddWord adds a word into the data structure
// (root A) -> [(root B) -> []]
func (root *LetterTrieNode) AddWord(word string) {
	if len(word) == 0 {
		return
	}
	for _, c := range word {
		node := root.findNode(c)
		if node == nil {
			node = makeNode(c)
			root.Children = append(root.Children, node)
		}
		root = node

	}
}

func makeNode(c rune) *LetterTrieNode {
	node := &LetterTrieNode{Children: make([]*LetterTrieNode, 0)}
	node.Letter = c
	return node
}

func (dict *WordDictionary) Search(word string) bool {
	if word == "" {
		return false
	}
	c, _ := utf8.DecodeRuneInString(word)
	fmt.Printf("%c\n", c)
	if c == '.' {
		for _, node := range dict.Children {
			if node.Search(word[1:]) {
				return true
			}
			return false
		}
	} else {
		root := findNode(dict.Children, c)
		fmt.Printf("root: %#v\n", root)
		return root.Search(word[1:])
	}
	return false
}

// Search returns if the word is in the data structure. A word could
// contain the dot character '.' to represent any one letter.
//
// Depth-first search implementation for simplicity
func (root *LetterTrieNode) Search(word string) bool {
	fmt.Printf("word2: %v\n", word)
	fmt.Printf("root2: %v\n", root)
	if root == nil {
		return false
	}
	fmt.Println("--START LOOP--")
	for i, c := range word {
		fmt.Println(c)
		if c == '.' {
			fmt.Println("wildcard baby")
			fmt.Println(root.Children)
			if len(root.Children) == 0 {
				return false
			}
			for _, node := range root.Children {
				if node.Search(word[i+1:]) {
					return true
				}
				return false
			}
		} else {
			node := root.findNode(c)
			fmt.Printf("node.Letter: %#v\n", node.Letter)
			fmt.Printf("c: %#v\n", c)
			if node == nil {
				return false
			} else if node.Letter == c {
				root = node
			} else {
				return false
			}
		}
	}
	return true
}

func (root *LetterTrieNode) findNode(c rune) *LetterTrieNode {
	if root != nil {
		node := findNode(root.Children, c)
		if node != nil {
			return node
		}
	}
	return nil
}

func findNode(a []*LetterTrieNode, c rune) *LetterTrieNode {
	for _, node := range a {
		if node.Letter == c {
			return node
		}
	}
	return nil
}

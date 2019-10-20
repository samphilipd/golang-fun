package main

import (
	"testing"
)

func TestLetterTrieNode(t *testing.T) {
	root := Constructor()

	root.AddWord("bad")
	root.AddWord("sad")
	root.AddWord("dad")
	// root.AddWord("mad")

	if root.Search("pad") != false {
		t.Fatal("BalancedStringSplit failed. Expected Search(\"pad\") to be false")
	}
	if root.Search("sad") != true {
		t.Fatal("BalancedStringSplit failed. Expected Search(\"sad\") to be true")
	}
	if root.Search("bad") != true {
		t.Fatal("BalancedStringSplit failed. Expected Search(\"bad\") to be true")
	}
	if root.Search(".ad") != true {
		t.Fatal("BalancedStringSplit failed. Expected Search(\".ad\") to be true")
	}
	if root.Search("b..") != true {
		t.Fatal("BalancedStringSplit failed. Expected Search(\"b..\") to be true")
	}
}

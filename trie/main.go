package main

import (
	"fmt"
)

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

func NewTrie() *TrieNode {
	return &TrieNode{}
}

func (tn *TrieNode) Insert(word string) {
	current := tn
	for _, ch := range word {
		index := int(ch) - 'a'
		if current.children[index] == nil {
			current.children[index] = NewTrie()
		}
		current = current.children[index]
	}
	current.isEnd = true
}

func (tn *TrieNode) Search(word string) bool {
	current := tn.findNode(word)
	if current == nil {
		return false
	}
	return current.isEnd
}

func (tn *TrieNode) StartsWith(prefix string) bool {
	current := tn.findNode(prefix)
	return current != nil
}

func (tn *TrieNode) GetWordsWithPrefix(prefix string) []string {
	words := make([]string, 0)
	if !tn.StartsWith(prefix) {
		return words
	}

	current := tn.findNode(prefix)
	if current == nil {
		return []string{}
	}
	return words
}
func dfs(node *TrieNode, current string, results *[]string) {

}

//dfs(node, wordSoFar, &results)
// if node.isEnd → append wordSoFar to results
// for i, child := range node.children
//   if child != nil
//     dfs(child, wordSoFar + string('a'+i), &results)

func (tn *TrieNode) Visualize() {
	fmt.Println("root")
	visualize(tn, "")
}

func (tn *TrieNode) findNode(prefix string) *TrieNode {
	current := tn
	for _, ch := range prefix {
		index := int(ch) - 'a'
		if current.children[index] == nil {
			return nil
		}
		current = current.children[index]
	}
	return current
}

func visualize(node *TrieNode, prefix string) {
	var childIndices []int
	for i := 0; i < 26; i++ {
		if node.children[i] != nil {
			childIndices = append(childIndices, i)
		}
	}
	for i, idx := range childIndices {
		isLast := i == len(childIndices)-1
		connector := "├── "
		continuation := "│   "
		if isLast {
			connector = "└── "
			continuation = "    "
		}
		label := string(rune('a' + idx))
		if node.children[idx].isEnd {
			label += " [*]"
		}
		fmt.Println(prefix + connector + label)
		visualize(node.children[idx], prefix+continuation)
	}
}

func main() {
	root := NewTrie()
	for _, word := range []string{"cat", "car", "card", "care", "bat"} {
		root.Insert(word)
	}
	//root.Visualize()

	fmt.Printf("%+v", root.GetWordsWithPrefix("car"))
}

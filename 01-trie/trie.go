package trie

import "encoding/json"

// NewNode 创建节点实例
func NewNode(b byte) *Node {
	return &Node{
		Data:     b,
		Children: make([]*Node, 26),
	}
}

// Node 节点
type Node struct {
	Data     byte
	Children []*Node
	IsEnding bool
}

// NewTrie 创建trie树实例
func NewTrie() *Trie {
	return &Trie{
		root: NewNode('/'),
	}
}

// Trie trie树
type Trie struct {
	root *Node
}

func (t *Trie) String() string {
	b, _ := json.Marshal(t.root)
	return string(b)
}

// Insert 插入数据
func (t *Trie) Insert(s []byte) {
	root := t.root

	for i := 0; i < len(s); i++ {
		idx := s[i] - 'a'
		if root.Children[idx] == nil {
			root.Children[idx] = NewNode(s[i])
		}
		root = root.Children[idx]
	}

	root.IsEnding = true
}

// Match 匹配字符串
func (t *Trie) Match(s []byte) bool {
	root := t.root

	for i := 0; i < len(s); i++ {
		idx := s[i] - 'a'
		if root.Children[idx] == nil {
			return false
		}
		root = root.Children[idx]
	}

	return root.IsEnding
}

// Find 查找字符串
func (t *Trie) Find(s []byte) []string {
	root := t.root

	var ss []string
	var prefix []byte
	for i := 0; i < len(s); i++ {
		idx := s[i] - 'a'
		if root.Children[idx] == nil {
			return nil
		}
		root = root.Children[idx]
		prefix = append(prefix, root.Data)
	}

	if root.IsEnding {
		ss = append(ss, string(prefix))
	}

	c := t.findChildren(prefix, root.Children)
	return append(ss, c...)
}

func (t *Trie) findChildren(prefix []byte, children []*Node) []string {
	var ss []string

	for i := 0; i < len(children); i++ {
		node := children[i]
		if node == nil {
			continue
		}

		prefix = append(prefix, node.Data)
		if node.IsEnding {
			ss = append(ss, string(prefix))
		}

		ss = append(ss, t.findChildren(prefix, node.Children)...)
	}

	return ss
}

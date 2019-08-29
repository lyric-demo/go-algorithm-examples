package trie

import "testing"

func TestMatch(t *testing.T) {
	trie := NewTrie()
	trie.Insert([]byte("hi"))
	trie.Insert([]byte("hel"))
	trie.Insert([]byte("hello"))
	trie.Insert([]byte("yes"))
	trie.Insert([]byte("yeah"))

	exists := trie.Match([]byte("hi"))
	if !exists {
		t.Error("不是期望的结果：", exists)
	}

	exists = trie.Match([]byte("hel"))
	if !exists {
		t.Error("不是期望的结果：", exists)
	}

	exists = trie.Match([]byte("hello"))
	if !exists {
		t.Error("不是期望的结果：", exists)
	}

	exists = trie.Match([]byte("hell"))
	if exists {
		t.Error("不是期望的结果：", exists)
	}
}

func TestFind(t *testing.T) {
	tr := NewTrie()

	tr.Insert([]byte("abc"))
	tr.Insert([]byte("abcd"))
	tr.Insert([]byte("abcde"))
	tr.Insert([]byte("abcdef"))
	tr.Insert([]byte("abcdefg"))
	tr.Insert([]byte("abcdefgh"))

	result := tr.Find([]byte("abcdef"))
	if len(result) != 3 ||
		result[0] != "abcdef" ||
		result[1] != "abcdefg" ||
		result[2] != "abcdefgh" {
		t.Error("Not expect result:", result)
		return
	}

	result = tr.Find([]byte("abc"))
	if len(result) != 6 {
		t.Error("Not expect result:", result)
		return
	}

}

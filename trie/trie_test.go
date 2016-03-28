package trie

import (
  "testing"
  "os"
)

func TestStrie(t *testing.T) {
	trie := NewStrie()
  testAddString(t, trie)
  testInsertion(t, trie)
}

func TestMtrie(t *testing.T) {
	trie := NewMtrie()
  testAddString(t, trie)
  testInsertion(t, trie)
}

func testAddString(t *testing.T, trie Trie) {
  trie.AddString("Foo")
  if trie.AddString("Foobar") { // shouldn't exist
    t.Error("AddString succeeded but shouldn't")
  }
  if !trie.AddString("Foo") { // should exist
    t.Error("AddString succeeded but shouldn't")
  }
}

func testInsertion(t *testing.T, trie Trie) {
  for i := 0; i < 2000000; i++ {
    r1 := rune(i%256)
    r2 := rune((i+1)%256)
    r3 := rune((i+2)%256)
    str := string([]rune{r1,r2,r3})
    trie.AddString(str)
  }
}

func TestPrintTrie(t *testing.T) {
  trie := NewMtrie()
  trie.AddString("D")
	trie.AddString("G")
  trie.AddString("H")
  trie.AddString("N")
  trie.AddString("P")
  trie.AddString("T")
  trie.AddString("T knows Python")
  trie.AddString("D knows TDD")
  trie.AddString("M knows Scala")
  trie.AddString("P knows Java")
  trie.AddString("G knows 🐐")
  trie.AddString("G")
  trie.PrintTrie(os.Stdout)
}

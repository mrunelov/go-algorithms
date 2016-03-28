package trie

/*
 * Trie implementations inspired by https://github.com/broersa/trie
 */

import (
  "fmt"
  "io"
)

var (
  // separator = " → "
  separator = ""
)

type Trie interface {
  AddString(s string) bool
  PrintTrie(w io.Writer)
  Contains(s string) bool
}

/*
 * A trie using slices internally
 */
type Strie struct {
  children []*Strie
  value    rune
  end      bool
}

/*
 * A trie using maps internally
 */
type Mtrie struct {
  children map[rune]*Mtrie
  end      bool
}

func NewMtrie() *Mtrie {
  return &Mtrie{
    children: make(map[rune]*Mtrie),
  }
}

func NewStrie() *Strie {
  return &Strie{
    children: make([]*Strie, 0),
  }
}

func (t *Strie) findNode(r rune) (*Strie, bool) {
  for _, child := range t.children {
    if child.value == r {
      return child, true
    }
  }
  return nil, false
}

func (t *Strie) Contains(str string) bool {
  // TODO
  return false
}

func (t *Mtrie) Contains(str string) bool {
  // TODO
  return false
}

/*
 * Adds a string to the trie.
 * returns true if the word already existed in the trie,
 * otherwise false
 */
func (t *Strie) AddString(str string) bool {
  collision := true
  current := t
  for _, r := range str {
    next, found := current.findNode(r)
    if !found {
      next = NewStrie()
      next.value = r
      current.children = append(current.children, next)
    }
    current = next
  }
  if !current.end {
    current.end = true
    collision = false
  }
  return collision
}

/*
 * Adds a string to the trie.
 * returns true if the word already existed in the trie,
 * otherwise false
 */
func (t *Mtrie) AddString(str string) bool {
  collision := true
  current := t
  for _, r := range str {
    next, found := current.children[r]
    if !found {
      next = NewMtrie()
      current.children[r] = next
    }
    current = next
  }
  if !current.end {
    current.end = true
    collision = false
  }
  return collision
}

func (t *Mtrie) PrintTrie(w io.Writer) {
  t.printWords(w, make([]rune, 0))
}

func (t *Strie) PrintTrie(w io.Writer) {
  t.printWords(w, make([]rune, 0))
}

func (t *Mtrie) printWords(w io.Writer, runes []rune) {
  if len(t.children) == 0 {
    fmt.Fprintf(w, "%v\n", string(runes))
  }
  for r := range t.children {
    child := t.children[r]
    word := append(runes, append([]rune(separator), r)...)
    if child.end {
      word = append(word, []rune("|")...)
    }
    child.printWords(w, word)
  }
}

func (t *Strie) printWords(w io.Writer, runes []rune) {
  if len(t.children) == 0 {
    fmt.Fprintf(w, "%v\n", string(runes))
  }
  for _, child := range t.children {
    word := append(runes, append([]rune(separator), child.value)...)
    if child.end {
      word = append(word, []rune("|")...)
    }
    child.printWords(w, word)
  }
}

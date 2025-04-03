package gwt

import (
	"testing"
)

func TestFullSearch(T *testing.T) {
	t := NewTrie()
	t.Insert("abcde")
	flag := t.FullSearch("abcd")
	if !flag {
		T.Errorf("full Search fail.....")
	}
}
func TestPrefixSearch(T *testing.T) {
	t := NewTrie()
	t.Insert("abcd")
	flag := t.PrefixSearch("ab")
	if !flag {
		T.Errorf("prefixSearch fail.......")
	}
}

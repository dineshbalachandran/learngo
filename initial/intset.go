package main

import (
	"bytes"
	"fmt"
)

//sz set to 32 or 64 depending on platform
const sz int = 32 << (^uint(0) >> 63)

type IntSet struct {
	words []uint
}

func (s *IntSet) Has(x int) bool {
	i, bit := x/sz, uint(x%sz)
	return i < len(s.words) && s.words[i]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	i, bit := x/sz, uint(x%sz)
	for i >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[i] |= 1 << bit
}

func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			break
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= (s.words[i] & tword)
		} else {
			break
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			break
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < sz; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", sz*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	len := 0
	for _, word := range s.words {
		for word > 0 {
			len += int(1 & word)
			word >>= 1
		}
	}
	return len
}

func (s *IntSet) Remove(x int) {
	i, bit := x/sz, uint(x%sz)
	if i < len(s.words) {
		s.words[i] ^= 1 << bit
	}
}

func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

func (s *IntSet) Copy() *IntSet {
	var t IntSet

	for _, word := range s.words {
		t.words = append(t.words, word)
	}

	return &t
}

func (s *IntSet) Elems() []int {
	var result []int
	for i, word := range s.words {
		for j := 0; j < sz; j++ {
			if word&(1<<uint(j)) != 0 {
				result = append(result, sz*i+j)
			}
		}
	}

	return result
}

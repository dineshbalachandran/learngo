package main

import (
	"sort"
)

func isPalindrome(s sort.Interface) bool {
	i := 0
	for i < s.Len()/2 {
		j := s.Len() - i - 1
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
		i++
	}

	return true
}

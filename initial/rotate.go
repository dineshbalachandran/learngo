package main

import "fmt"

func rotate(n int, a []int) {

	for _, v := range a {
		fmt.Printf("%d", v)
	}
	l := len(a)
	r := n % l

	s := a[l-r:]
	for _, v := range a[:l-r] {
		s = append(s, v)
	}

	for _, v := range s {
		fmt.Printf("%d", v)
	}
}

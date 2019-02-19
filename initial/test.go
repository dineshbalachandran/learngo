package main

import (
	"fmt"
	"strings"
)

func test() {

	var z []int

	z = append(z, 1)
	x := append(z, 2)
	x = append(z, 3)
	test1(&x)

	fmt.Println(cap(z))
	fmt.Println(cap(x))

	for _, v := range z {
		fmt.Println(v)
	}

	for _, v := range x {
		fmt.Println(v)
	}

}

func test1(a *[]int) {
	*a = append(*a, 5)
}

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "$foo", f("foo"), -1)
}

type s struct {
	s1 int
	s2 int
}

func (p s) String() string {
	return fmt.Sprintf("%d %d", p.s1, p.s2)
}

func testmap() {
	m := make(map[string]*s)
	s1 := s{s1: 1, s2: 2}

	m["1"] = &s1
	s1.s1 = 3

	fmt.Println(m["1"])
}

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

//rotate in a single pass inline using no other data structure
func onepassrotate(n int, a []int) {

	for _, v := range a {
		fmt.Printf("%d", v)
	}

	l := len(a)
	r := n % l

	fmt.Println()
	fmt.Println(r)

	i, k := 1, 1
	v := a[i-1]
	for k <= l {
		i += r
		if i > l {
			i -= l
		}

		a[i-1], v = v, a[i-1]

		//special handling is required if the rotation 'r' is a factor of the count 'l'
		//(i.e. when l%r==0)
		if l%r == 0 && k%(l/r) == 0 {
			i++
			v = a[i-1]
		}

		k++
	}

	for _, v := range a {
		fmt.Printf("%d", v)
	}
}

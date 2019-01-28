package main

import "fmt"

func test() {

	var z []int

	z = append(z, 1)
	x := append(z, 2)
	x = append(z, 3)

	for i, v := range z {
		fmt.Println(v)
		fmt.Println(x[i])
	}

}

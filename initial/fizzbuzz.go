package main

import (
	"fmt"
)

func fizzbuzz() {

	for i := 1; i <= 15; i++ {
		var s string
		if i%3 == 0 {
			s += "Fizz"
		}
		if i%5 == 0 {
			s += "Buzz"
		}

		if len(s) > 0 {
			fmt.Println(s)
		} else {
			fmt.Println(i)
		}
	}
}

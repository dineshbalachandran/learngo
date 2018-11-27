package main

import (
	"fmt"
	"github.com/dineshbalachandran/shapes"
)

func main() {

	printarea()
	fizzbuzz()
}

func printarea() {
	shapes := []shapes.Shape{shapes.Circle{X: 0.0, Y: 0.0, R: 5.0},
		shapes.Rectangle{X1: 0, Y1: 0, X2: 2, Y2: 2}}

	for _, s := range shapes {
		fmt.Println(s.Area())
	}
}

func fizzbuzz() {

	for i := 1; i <= 100; i++ {
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

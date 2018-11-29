package main

import (
	"fmt"
	"github.com/dineshbalachandran/shapes"
	"os"
)

func main() {
	dup()
}

func printargs() {
	for i, a := range os.Args {
		fmt.Println(i, a)
	}
}

func printarea() {
	shapes := []shapes.Shape{shapes.Circle{X: 0.0, Y: 0.0, R: 5.0},
		shapes.Rectangle{X1: 0, Y1: 0, X2: 2, Y2: 2}}

	for _, s := range shapes {
		fmt.Println(s.Area())
	}
}

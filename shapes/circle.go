package shapes

import "math"

type Circle struct {
	X, Y, R float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.R * c.R
}

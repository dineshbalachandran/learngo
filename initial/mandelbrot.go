package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

// MandelbrotImage generates a mandelbrot image
func MandelbrotImage() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const (
		iterations = 200
		contrast   = 15
	)

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			//			var r, g, b uint8
			//			switch {
			//			case n <= 50:
			//				r, g, b = n, 0, 0
			//			case n <= 100:
			//				r, g, b = 0, n, 0
			//			case n <= 150:
			//				r, g, b = 0, 0, n
			//			case n <= 200:
			//				r, g, b = 0, 0, 0
			//			}
			return color.RGBA{255 - contrast*n, 255 - contrast*n, 255 - contrast*n, 255}
		}
	}

	return color.Black
}

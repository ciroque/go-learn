package main

import (
	"golang.org/x/tour/pic"
	"math/rand"
)

func Pic(dx, dy int) [][]uint8 {
	xs := make([][]uint8, dx)

	var funcs []func(x, y int) uint8
	funcs = append(funcs, gradient)
	funcs = append(funcs, 	napkin)
	funcs = append(funcs, mosaic)
	funcs = append(funcs, ovalers)
	funcs = append(funcs, cartesianPlane)
	funcs = append(funcs, circulars)

	function := funcs[rand.Intn(len(funcs))]

	for ix := range xs {
		ys := make([]uint8, dy)
		for iy := range ys {
			ys[iy] = function(ix, iy)
		}
		xs[ix] = ys
	}
	return xs
}

func gradient(x, y int) uint8 {
	return uint8((x + y) / 2)
}

func napkin(x, y int) uint8 {
	return uint8(-x ^ y)
}

func mosaic(x, y int) uint8 {
	return uint8((x ^ (x * y)) + (y ^ (y * x)))
}

func ovalers(x, y int) uint8 {
	return uint8((x * x/1) + (y * y/1) / 7)
}

func cartesianPlane(x, y int) uint8 {
	v := (x ^ (x + 1)) * (y ^ (y + 1))
	return uint8( v )
}

func circulars(x, y int) uint8 {
	divisor := 5 // play with this to offset the circles
	return uint8((x * x/divisor) + (y * y/divisor))
}

func main() {
	pic.Show(Pic)
}

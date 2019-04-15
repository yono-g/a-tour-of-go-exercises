package main

import (
	"fmt"
	"math"
)

const delta = 0.000001

func Sqrt(x float64) float64 {
	z := x / 2
	for loop := 0; ; loop++ {
		y := z - (z*z-x)/(2*z)
		if y == z || math.Abs(y-z) < delta {
			fmt.Println("loop count:", loop)
			break
		}

		z = y
	}

	return z
}

func main() {
	x := 227.0
	fmt.Println(Sqrt(x))
	fmt.Println(math.Sqrt(x))
}

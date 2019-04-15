package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number", float64(e)) // Sprintがeをstringに変換しようとして、Stringerが満たされていないのでError()が呼び出され、無限に呼び出されると思われる
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	const delta = 0.00001

	z := x / 2
	for loop := 0; ; loop++ {
		y := z - (z*z-x)/(2*z)
		if y == z || math.Abs(y-z) < delta {
			break
		}

		z = y
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}


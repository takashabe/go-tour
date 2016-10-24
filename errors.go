// https://go-tour-jp.appspot.com/methods/20

package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return "cannot Sqrt negative number: " + fmt.Sprint(float64(e))
}

func Sqrt(x float64) (float64, error) {
	// error check
	if x < 0 {
		e := ErrNegativeSqrt(x)
		return x, e
	}

	z := 1.0
	var tmp float64
	for z != tmp {
		tmp = z
		z = (z + x/z) / 2
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

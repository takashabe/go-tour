// https://go-tour-jp.appspot.com/flowcontrol/8

package main

import (
	"fmt"
)

// Use algorithm "Newton's method". => https://ja.wikipedia.org/wiki/%E3%83%8B%E3%83%A5%E3%83%BC%E3%83%88%E3%83%B3%E6%B3%95
func Sqrt(x float64) float64 {
	z := 1.0
	var tmp float64
	for z != tmp {
		tmp = z
		z = (z + x/z) / 2
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}

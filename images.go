// https://go-tour-jp.appspot.com/methods/25

package main

import (
	"golang.org/x/tour/pic"
	"image/color"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
}

func (i Image) Bounds() Rectangle {
}

func (i Image) At(x, y int) color.Color {
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}

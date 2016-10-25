// https://go-tour-jp.appspot.com/methods/23

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	for i, v := range p {
		if ('A' <= v && v <= 'M') || ('a' <= v && v <= 'm') {
			p[i] += 13
		} else if ('N' <= v && v <= 'Z') || ('n' <= v && v <= 'z') {
			p[i] -= 13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

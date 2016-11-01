// https://go-tour-jp.appspot.com/concurrency/7
// https://go-tour-jp.appspot.com/concurrency/8

package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	fmt.Println(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	// TODO change me
	return false
}

func main() {
	t := tree.New(2)
	c := make(chan int)
	Walk(t, c)
}

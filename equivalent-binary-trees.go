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
	_walk(t, ch)
	close(ch)
}

func _walk(t *tree.Tree, ch chan int) {
	if t != nil {
		_walk(t.Left, ch)
		ch <- t.Value
		_walk(t.Right, ch)
	}
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
	go Walk(t, c)
	for i := range c {
		fmt.Println(i)
	}
}

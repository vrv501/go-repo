package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch := make(chan int)
	ch2 := make(chan int)
	go func() {
		Walk(t1, ch)
		close(ch)
	}()

	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	for {
		val, ok := <-ch
		val2, ok2 := <-ch2
		if ok != ok2 || val != val2 {
			return false
		}

		if ok {
			break
		}
	}
	return true
}

func main() {

	fmt.Println(Same(tree.New(1), tree.New(1)))
}

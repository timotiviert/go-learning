package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)

	var walk func(t *tree.Tree)
	walk = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walk(t.Left)
		ch <- t.Value
		walk(t.Right)
	}
	walk(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := make(chan int), make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)

	// While true loop.
	for {
		// Read values from channels. ok1 and ok2 false when channel closed.
		v1, ok1 := <-c1
		v2, ok2 := <-c2

		// Values of t1 and t2 are different or different size (one channel is closed and other still open).
		if v1 != v2 || ok1 != ok2 {
			return false
		}
		// Channel 1 closed and as condition before was not true Channel 2 as well. Therefore Walks are finished.
		if !ok1 {
			break
		}
	}
	return true
}

func main() {
	// Test Walk
	c := make(chan int)
	t1 := tree.New(1)
	go Walk(t1, c)
	// Range zieht automatisch Wert aus c.. kein <-c nötig.
	for v := range c {
		println(v)
	}

	t2 := tree.New(2)
	fmt.Println(Same(t1, tree.New(1)))
	fmt.Println(Same(t1, t2))
}

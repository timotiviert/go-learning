package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n := -1
	p := 1
	pp := 0

	return func() int {
		n++
		if n == 0 {
			return 0
		} else if n == 1 {
			return 1
		}
		fib := pp + p
		pp = p
		p = fib
		return fib
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}

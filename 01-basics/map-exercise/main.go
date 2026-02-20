package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	count := make(map[string]int)
	words := strings.Fields(s)

	//for _, word := range words {
	//	elem, ok := count[word]
	//
	//	if ok == true {
	//		count[word] = elem + 1
	//	} else {
	//		count[word] = 1
	//	}
	//}
	for _, word := range words {
		count[word]++
	}

	return count
}

func main() {
	wc.Test(WordCount)
}

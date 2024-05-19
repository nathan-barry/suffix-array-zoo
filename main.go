package main

import (
	"fmt"
)

func main() {
	text := "banana"

	SA := naiveSolution(text)

	printSuffixArray(text, SA)
}

func printSuffixArray(text string, SA []int) {
	for _, i := range SA {
		fmt.Printf("%v : %v\n", i, text[i:])
	}
}

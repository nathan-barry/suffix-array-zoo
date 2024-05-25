package main

import "fmt"

func main() {
	text := loadBytesFromTextFile("data/man-in-the-arena.txt")

	SA := NewSuffixArray(text, NAIVE)

	fmt.Printf("Arena: contains %v, index %v\n", SA.Contains("arena"), SA.Find("arena"))
	fmt.Printf("Poop: contains %v, index %v\n", SA.Contains("poop"), SA.Find("poop"))
}

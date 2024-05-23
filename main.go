package main

import "fmt"

func main() {
	text := loadFromTextFile("data/man-in-the-arena.txt")

	SA := NewSuffixArray(text)

	naiveCA(SA)

	fmt.Printf("Arena: contains %v, index %v\n", SA.Contains("arena"), SA.Find("arena"))
	fmt.Printf("Poop: contains %v, index %v\n", SA.Contains("poop"), SA.Find("poop"))
}

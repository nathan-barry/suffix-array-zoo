package main

import "fmt"

func main() {
	// data := loadBytesFromTextFile("data/man-in-the-arena.txt")
	data := []byte("processing")
	saca := DC3_SKEW // Suffix Array Construction Algorithm

	sa := NewSuffixArray(data, saca)

	fmt.Printf("Arena: contains %v, index %v\n", sa.Contains("arena"), sa.Find("arena"))
	fmt.Printf("Poop: contains %v, index %v\n", sa.Contains("poop"), sa.Find("poop"))
}

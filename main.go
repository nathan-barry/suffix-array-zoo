package main

import "fmt"

func main() {
	// data := loadBytesFromTextFile("data/man-in-the-arena.txt")
	data := []byte("nathan is awesome yabadabadoo")
	saca := DC3_SKEW // Suffix Array Construction Algorithm

	sa := NewSuffixArray(data, saca)
	fmt.Println("dc3")
	sa.Print()

	saca2 := NAIVE
	fmt.Println("\n\nnaive")
	sa2 := NewSuffixArray(data, saca2)
	sa2.Print()
}

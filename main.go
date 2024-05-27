package main

func main() {
	// data := loadBytesFromTextFile("data/man-in-the-arena.txt")
	data := []byte("processing")
	saca := DC3_SKEW // Suffix Array Construction Algorithm

	sa := NewSuffixArray(data, saca)

	sa.Print()
}

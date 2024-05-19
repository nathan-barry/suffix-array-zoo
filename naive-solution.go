package main

import "sort"

func naiveSolution(text string) []int {
	// Initialize suffix array
	suffixArray := make([]int, len(text))
	for i := range suffixArray {
		suffixArray[i] = i
	}

	// Sort the indexes based on the suffixes
	sort.Slice(suffixArray, func(i, j int) bool {
		return text[suffixArray[i]:] < text[suffixArray[j]:]
	})

	return suffixArray
}

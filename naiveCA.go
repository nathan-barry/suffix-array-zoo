package main

import "sort"

func naiveCA(SA SuffixArray) {
	// Initialize suffix array
	for i := range SA.indices {
		SA.indices[i] = i
	}

	// Sort indices in lexicographical order
	sort.Slice(SA.indices, func(i, j int) bool {
		return (*SA.text)[SA.indices[i]:] < (*SA.text)[SA.indices[j]:]
	})
}

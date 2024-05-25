package main

import "sort"

func (sa *SuffixArray) naive() {
	// Initialize suffix array
	for i := range sa.sa {
		sa.sa[i] = i
	}

	text := string(sa.data)

	// Sort sa in lexicographical order
	sort.Slice(sa.sa, func(i, j int) bool {
		return text[sa.sa[i]:] < text[sa.sa[j]:]
	})
}

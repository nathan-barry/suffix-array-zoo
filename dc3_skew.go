package main

import "sort"

func dc3Skew(SA SuffixArray) {
	// TODO
	// 1. Split text into S1, S2 triples, concat for S12
	//	- Concat them to make S12 (with indices?)
	//	- Radix sort them
	//	- If no repeats, return SA12
	//	- Else, run again with numbers as string
	// 2. Sort S0 with S12
	//	- End with sorted S0 and sorted S12
	// 3. Merge S0 and S12
	//	- Start with left of S0 and S12, compare and insert into
	//	  final SA with trick (look at gist notes for how)

	// Initialize suffix array
	for i := range SA.indices {
		SA.indices[i] = i
	}

	// Sort indices in lexicographical order
	sort.Slice(SA.indices, func(i, j int) bool {
		return (*SA.text)[SA.indices[i]:] < (*SA.text)[SA.indices[j]:]
	})
}

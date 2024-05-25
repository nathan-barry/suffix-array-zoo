package main

func (sa *SuffixArray) dc3Skew() {
	// Append 3 sentinel end characters
	for i := 0; i < 3; i++ {
		sa.sa = append(sa.sa, 0)
	}

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

}

// %%%%%%%%%% Helpers %%%%%%%%%%

// Stably sort a[0..n-1] to b[0..n-1] with keys in 0..K from r
func radixPass(a, b, r []int, n, K int) {

}

// lexicographic order for pairs
func leq2(a1, a2, b1, b2 int) bool {
	return (a1 < b1) || (a1 == b1 && a2 <= b2)
}

// lexicographic order for triples
func leq3(a1, a2, a3, b1, b2, b3 int) bool {
	return (a1 < b1) || (a1 == b1 && leq2(a2, a3, b2, b3))
}

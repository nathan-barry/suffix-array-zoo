package main

import "fmt"

const K = 256 // alphabet is a byte

func (sa *SuffixArray) dc3Skew() {
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
	sa.data = append(sa.data, 0)
	sa.data = append(sa.data, 0)
	sa.data = append(sa.data, 0)

	dc3SkewRecurse(sa.data, sa.sa, len(sa.data)-3)
}

// Ex "abcabca"
// n = 7
// n0 = (a a a) => 3
// n1 = (b b) => 2
// n2 = (c c) => 2
// n12 = (b b c c) => 4
// triples12 = (bca bca cab ca\0)
func dc3SkewRecurse(s []byte, sa []int, n int) {
	n0 := (n + 2) / 3
	n1 := (n + 1) / 3
	n2 := n / 3
	n12 := n1 + n2

	s12 := make([]byte, n12*3)

	// Debug information
	fmt.Println(n, n0, n1, n2)
	fmt.Println(n, string(s))
	fmt.Println(n, s)
	printGroups(s, n0, n1, n2)

	gen3grams(s, s12, n1, n2)

	// construct s12
	// sa12 := make([]int, n12)

	// i0 = 0, 3, 6...
	// i1 = 1, 4, 7...
	// i2 = 2, 5, 8...

	fmt.Println(n12, string(s12))
	fmt.Println(n12, s12)
}

// %%%%%%%%%% Helpers %%%%%%%%%%

// SA[rank] = pos
// ISA[pos] = rank

func gen3grams(s, s12 []byte, n1, n2 int) {
	// Add 3-gram from s1
	for i := 0; i < n1; i++ {
		fmt.Println(i, s[1+(i*3):1+(i*3)+3], 1+(i*3), 1+(i*3)+3)
		copy(s12[i*3:(i*3)+3], s[1+(i*3):1+(i*3)+3])
	}

	// // Add 3-gram from s1
	for i := 0; i < n2; i++ {
		fmt.Println(i, s[2+(i*3):2+(i*3)+3], 2+(i*3), 2+(i*3)+3)
		copy(s12[(n1+i)*3:((n1+i)*3)+3], s[2+(i*3):2+(i*3)+3])
	}
}

func printGroups(s []byte, n0, n1, n2 int) {
	s0 := make([]byte, n0*3)
	s1 := make([]byte, n1*3)
	s2 := make([]byte, n2*3)

	for i := 0; i < n0; i++ {
		s0[i] = s[i*3]
	}
	fmt.Println(n0, string(s0))

	for i := 0; i < n1; i++ {
		s1[i] = s[1+i*3]
	}
	fmt.Println(n1, string(s1))

	for i := 0; i < n2; i++ {
		s2[i] = s[2+i*3]
	}
	fmt.Println(n2, string(s2))
}

// Stably sort a[0..n-1] to b[0..n-1] with keys in 0..K from r
func radixPass(a, b, r []int, n int) {
}

// lexicographic order for pairs
func leq2(a1, a2, b1, b2 int) bool {
	return (a1 < b1) || (a1 == b1 && a2 <= b2)
}

// lexicographic order for triples
func leq3(a1, a2, a3, b1, b2, b3 int) bool {
	return (a1 < b1) || (a1 == b1 && leq2(a2, a3, b2, b3))
}

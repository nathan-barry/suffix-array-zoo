package main

import (
	"fmt"
	"log"
)

const K = 256 // alphabet is a byte

type gram [3]byte

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

func dc3SkewRecurse(s []byte, sa []int, n int) {
	n0 := (n + 2) / 3
	n1 := (n + 1) / 3
	n2 := n / 3

	n12 := n1 + n2           // number of bytes
	r12 := make([]int, n12)  // r12[pos in s] = rank
	sa12 := make([]int, n12) // sa12[rank] = pos in s

	// Step 0: generate positions of mod 1 and mod 2 suffixes
	for i, j := 0, 0; i < n; i++ {
		if i%3 != 0 {
			r12[j] = i
			j++
		}
	}

	fmt.Println(n, n0, n1, n2)
	fmt.Println(n12, r12)
	fmt.Println(n12, sa12)

	fmt.Printf("byte 3-grams: ")
	for _, i := range r12 {
		fmt.Printf("%v ", s[i:i+3])
	}
	fmt.Print("\n")

	// lsb radix sort s12 3-grams
	radixPass(s, r12, sa12, n12, 2) // r12 temp suffix positions, sa12 temp storage
	fmt.Println(n12, sa12)

	radixPass(s, sa12, r12, n12, 1) // sa12 temp sorted positions, r12 temp storage
	radixPass(s, r12, sa12, n12, 0) // r12 temp sorted positions, sa12 temp storage

	fmt.Println(n12, sa12)

	log.Fatal("end")

	// Debug information
	fmt.Println(n, n0, n1, n2)
	fmt.Println(n, string(s))
	fmt.Println(n, s)

}

// SA[rank] = pos
// ISA[pos] = rank

// %%%%%%%%%% Helpers %%%%%%%%%%

// `in`: input array containing the positions of the suffixes to be sorted
// `out`: output array where the sorted positions will be sorted
// `kth`: passed as 2, 1, then 0
// Essentially one pass of counting sort
func radixPass(s []byte, in, out []int, n12, kth int) {
	c := make([]int, K+1)

	// reset counts
	for i := 0; i <= K; i++ {
		c[i] = 0
	}

	// count occurrences of byte char
	for i := 0; i < n12; i++ {
		// `in[i]+kth` is suffix start+kth position
		// `s[^]` is the byte at that position in s
		// ^ byte is seen and the count incremented
		c[s[in[i]+kth]]++
	}
	fmt.Println("count c:", c)

	// turn c into exclusive prefix sums
	for i, sum := 0, 0; i <= K; i++ {
		temp := c[i]
		c[i] = sum
		sum += temp // equals n12 at end
	}

	fmt.Println("prefix sum c:", c)
	// store sorted positions
	fmt.Println("0 out:", out)
	for i := 0; i < n12; i++ {
		// `s[in[i]+kth]` is the kth byte at that suffix position
		// c[^] is the prefix sum of that byte
		fmt.Println("s[in[i]+kth]", s[in[i]+kth])
		prefix_sum := c[s[in[i]+kth]]
		fmt.Println("prefix_sum, in[i]", prefix_sum, in[i])
		out[prefix_sum] = in[i]
		// increment (part of counting sort)
		c[s[in[i]+kth]]++
		fmt.Println(i+1, "out:", out)
	}
}

func print3Grams(s []byte, indices []int) {
	for _, i := range indices {
		fmt.Printf("%3d ", i)
	}
	fmt.Print("\n")
	for _, i := range indices {
		fmt.Printf("%v ", string(s[i:i+3]))
	}
	fmt.Print("\n")
}

// func gen3grams(s, s12 []byte, n1, n2 int) {
// 	// Add 3-grams from s1
// 	for i := 0; i < n1; i++ {
// 		// fmt.Println(i, s[1+(i*3):1+(i*3)+3], 1+(i*3), 1+(i*3)+3)
// 		copy(s12[i*3:(i*3)+3], s[1+(i*3):1+(i*3)+3])
// 	}

// 	// // Add 3-grams from s2
// 	for i := 0; i < n2; i++ {
// 		// fmt.Println(i, s[2+(i*3):2+(i*3)+3], 2+(i*3), 2+(i*3)+3)
// 		copy(s12[(n1+i)*3:((n1+i)*3)+3], s[2+(i*3):2+(i*3)+3])
// 	}
// }

// func printGroups(s []byte, n0, n1, n2 int) {
// 	s0 := make([]byte, n0*3)
// 	s1 := make([]byte, n1*3)
// 	s2 := make([]byte, n2*3)

// 	for i := 0; i < n0; i++ {
// 		s0[i] = s[i*3]
// 	}
// 	fmt.Println(n0, string(s0))

// 	for i := 0; i < n1; i++ {
// 		s1[i] = s[1+i*3]
// 	}
// 	fmt.Println(n1, string(s1))

// 	for i := 0; i < n2; i++ {
// 		s2[i] = s[2+i*3]
// 	}
// 	fmt.Println(n2, string(s2))
// }

// lexicographic order for pairs
func leq2(a1, a2, b1, b2 int) bool {
	return (a1 < b1) || (a1 == b1 && a2 <= b2)
}

// lexicographic order for triples
func leq3(a1, a2, a3, b1, b2, b3 int) bool {
	return (a1 < b1) || (a1 == b1 && leq2(a2, a3, b2, b3))
}

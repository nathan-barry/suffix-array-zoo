package main

import (
	"fmt"
	"log"
)

const BYTE_SIZE = 256

// ALGORITHM OVERVIEW (high level, not exact implementation)
// 1. Split text into s1, s2 3-grams, concat for s12
//   - Concat them to make s12
//   - Radix sort them
//   - If no repeats, return sorted sa12
//   - else, run again with r12 as string
//
// 2. Induce sa0 from s0 and sa12
//   - End with sorted sa0 and sorted sa12
//
// 3. Merge sa0 and sa12
//   - Start with left of sa0 and sa12, compare and insert into final sa
func (sa *SuffixArray) dc3Skew() {
	// add sentinel padding to avoid out-of-bounds indexing
	sa.data = append(sa.data, 0)
	sa.data = append(sa.data, 0)
	sa.data = append(sa.data, 0)

	dc3SkewRecurse(sa.data, sa.sa, len(sa.data)-3, BYTE_SIZE)
}

func dc3SkewRecurse(s []byte, sa []int, n, K int) {
	n0 := (n + 2) / 3
	n1 := (n + 1) / 3
	n2 := n / 3

	n12 := n1 + n2              // number of bytes
	r12 := make([]int, n12)     // r12[pos in s12] = rank
	sa12 := make([]int, n12)    // sa12[rank-1] = pos in s12
	name12 := make([]byte, n12) // same as r12 but byte instead

	// Step 0: generate positions of mod 1 and mod 2 suffixes
	// temporarily store in r12 for memory efficiency
	for i, j := 0, 0; i < n; i++ {
		if i%3 != 0 {
			r12[j] = i
			j++
		}
	}

	// Step 1: Sort sa12
	// lsb radix sort s12 3-grams
	radixPass(s, r12, sa12, n12, 2, K) // r12 temp suffix positions, sa12 temp storage
	radixPass(s, sa12, r12, n12, 1, K) // sa12 temp sorted positions, r12 temp storage
	radixPass(s, r12, sa12, n12, 0, K) // r12 temp sorted positions, sa12 store output

	// find lexicographical names of 3-grams and write them to correct place in r12
	name, prev_c0, prev_c1, prev_c2 := 0, -1, -1, -1
	for i := 0; i < n12; i++ {
		cur_c0 := int(s[sa12[i]])
		cur_c1 := int(s[sa12[i]+1])
		cur_c2 := int(s[sa12[i]+2])
		// if current suffix is different from last suffix, update
		if cur_c0 != prev_c0 || cur_c1 != prev_c1 || cur_c2 != prev_c2 {
			name++
			s[sa12[i]] = byte(cur_c0)
			s[sa12[i]+1] = byte(cur_c1)
			s[sa12[i]+2] = byte(cur_c2)
		}
		// below mimics r1 concat r2
		if sa12[i]%3 == 1 { // from s1 group
			r12[sa12[i]/3] = name
			name12[sa12[i]/3] = byte(name)
		} else { // from s2 group
			r12[sa12[i]/3+n2] = name
			name12[sa12[i]/3] = byte(name)
		}
	}

	// recurse if names are not yet unique
	if name < n12 {
		dc3SkewRecurse(name12, sa12, n12, name)
		// store unique names in r12 using the suffix array
		for i := 0; i < n12; i++ {
			r12[sa12[i]] = i + 1 // r12 starts at 1, not 0
		}
	} else {
		// generate sa12 from r12 directly
		// sa12 currently holds index of suffix in s, not s12
		for i := 0; i < n12; i++ {
			sa12[r12[i]-1] = i // r12 starts at 1, not 0
		}
	}

	// Step 2: Sort sa0 from sa12
	r0 := make([]int, n0)  // r12[pos in s12] = rank
	sa0 := make([]int, n0) // sa12[rank-1] = pos in s12

	fmt.Printf("n0: %v, n1: %v, n2: %v, n12: %v\n", n0, n1, n2, n12)
	fmt.Printf("r12: %v, sa12: %v, r0: %v, sa0: %v\n", r12, sa12, r0, sa0)

	// stably sort the i mod 3 == 0 suffixes by rank of the i+1 suffixes
	for i, j := 0, 0; i < n12; i++ {
		if sa12[i] < n0 { // note: if (i == n0) > n1, no edge case since end padding?
			r0[j] = 3 * sa12[i]
			j++
		}
	}
	// sort by the first character in suffix, gives us sorted sa0 with s index position
	radixPass(s, r0, sa0, n0, 0, K)

	fmt.Printf("n0: %v, n1: %v, n2: %v, n12: %v\n", n0, n1, n2, n12)
	fmt.Printf("r12: %v, sa12: %v, r0: %v, sa0: %v\n", r12, sa12, r0, sa0)

	// Step 3: Merge
	// merge sorted sa0 suffixes and sorted sa12 suffixes
	for p, t, k := 0, n0-n1, 0; k < n; k++ {

	}

	log.Fatal("end")

}

// %%%%%%%%%% Helpers %%%%%%%%%%

// `in`: input array containing the positions of the suffixes to be sorted
// `out`: output array where the sorted positions will be sorted
// `kth`: passed as 2, 1, then 0
// Essentially one pass of counting sort
func radixPass(s []byte, in, out []int, n12, kth, K int) {
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

	// turn c into exclusive prefix sums
	for i, sum := 0, 0; i <= K; i++ {
		temp := c[i]
		c[i] = sum
		sum += temp // equals n12 at end
	}

	// store sorted positions
	for i := 0; i < n12; i++ {
		// `s[in[i]+kth]` is the kth byte at that suffix position
		// c[^] is the prefix sum of that byte
		prefix_sum := c[s[in[i]+kth]]
		out[prefix_sum] = in[i]
		// increment (part of counting sort)
		c[s[in[i]+kth]]++
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

// lexicographic order for pairs
func leq2(a1, a2, b1, b2 int) bool {
	return (a1 < b1) || (a1 == b1 && a2 <= b2)
}

// lexicographic order for triples
func leq3(a1, a2, a3, b1, b2, b3 int) bool {
	return (a1 < b1) || (a1 == b1 && leq2(a2, a3, b2, b3))
}

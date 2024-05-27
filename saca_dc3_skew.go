package main

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

	n02 := n0 + n2                // why n02 instead of n02, we'll never know
	r12 := make([]int, n02+3)     // r12[pos in s12] = rank
	sa12 := make([]int, n02+3)    // sa12[rank-1] = pos in s12
	name12 := make([]byte, n02+3) // same as r12 but byte instead

	// Step 0: generate positions of mod 1 and mod 2 suffixes
	// temporarily store in r12 for memory efficiency
	for i, j := 0, 0; i < n+(n0-n1); i++ {
		if i%3 != 0 {
			r12[j] = i
			j++
		}
	}

	// Step 1: Sort sa12
	// lsb radix sort s12 3-grams
	radixPass(s, r12, sa12, n02, 2, K) // r12 temp suffix positions, sa12 temp storage
	radixPass(s, sa12, r12, n02, 1, K) // sa12 temp sorted positions, r12 temp storage
	radixPass(s, r12, sa12, n02, 0, K) // r12 temp sorted positions, sa12 store output

	// find lexicographical names of 3-grams and write them to correct place in r12
	name, prev_c0, prev_c1, prev_c2 := 0, -1, -1, -1
	for i := 0; i < n02; i++ {
		cur_c0 := int(s[sa12[i]])
		cur_c1 := int(s[sa12[i]+1])
		cur_c2 := int(s[sa12[i]+2])
		// if current suffix is different from last suffix, update
		if cur_c0 != prev_c0 || cur_c1 != prev_c1 || cur_c2 != prev_c2 {
			name++
			prev_c0 = cur_c0
			prev_c1 = cur_c1
			prev_c2 = cur_c2
		}
		// below mimics r1 concat r2
		if sa12[i]%3 == 1 { // from s1 group
			r12[sa12[i]/3] = name
			name12[sa12[i]/3] = byte(name)
		} else { // from s2 group
			r12[sa12[i]/3+n0] = name
			name12[sa12[i]/3+n0] = byte(name)
		}
	}

	// recurse if names are not yet unique
	if name < n02 {
		dc3SkewRecurse(name12, sa12, n02, name)
		// store unique names in r12 using the suffix array
		for i := 0; i < n02; i++ {
			r12[sa12[i]] = i + 1 // r12 starts at 1, not 0
		}
	} else {
		// generate sa12 from r12 directly
		// sa12 currently holds index of suffix in s, not s12
		for i := 0; i < n02; i++ {
			sa12[r12[i]-1] = i // r12 starts at 1, not 0
		}
	}

	// Step 2: Sort sa0 from sa12
	r0 := make([]int, n0)  // r12[pos in s12] = rank
	sa0 := make([]int, n0) // sa12[rank-1] = pos in s12

	// stably sort the i mod 3 == 0 suffixes by rank of the i+1 suffixes
	for i, j := 0, 0; i < n02; i++ {
		if sa12[i] < n0 { // note: if (i == n0) > n1, no edge case since end padding?
			r0[j] = 3 * sa12[i]
			j++
		}
	}
	// sort by the first character in suffix, gives us sorted sa0 with s index position
	radixPass(s, r0, sa0, n0, 0, K)

	// Step 3: Merge
	// merge sorted sa0 suffixes and sorted sa12 suffixes
	for k, p0, p12 := 0, 0, 0; k < n; k++ {
		// `k` is the index iterating through sa
		// `p0` is the index iterating through sa0
		// `p12` is the index iterating through sa12

		i12 := getSA12Index(sa12, n0, p12) // pos of current offset s12 suffix in s
		i0 := sa0[p0]                      // pos of current offset s0 suffix in s

		if sa12Smaller(s, sa12, r12, n0, p12, i12, i0) {
			// suffix from sa12 is smaller
			sa[k] = i12
			p12++
			if p12 == n02 { // done, only sa0 suffixes left
				k++
				for p0 < n0 {
					sa[k] = sa0[p0]
					k++
					p0++
				}
			}
		} else {
			// suffix from sa0 is smaller
			sa[k] = i0
			p0++
			if p0 == n0 { // done, only sa12 suffixes left
				k++
				for p12 < n02 {
					sa[k] = getSA12Index(sa12, n0, p12)
					k++
					p12++
				}
			}
		}
	}
}

func getSA12Index(sa12 []int, n0, p12 int) int {
	if sa12[p12] < n0 { // suffix is mod 1
		return sa12[p12]*3 + 1
	} else { // suffix is mod 2
		return (sa12[p12]-n0)*3 + 2
	}
}

func sa12Smaller(s []byte, sa12, r12 []int, n0, p12, i12, i0 int) bool {
	if sa12[p12] < n0 { // different compares for mod 1 and mod 2 suffixes
		return leq2(int(s[i12]), r12[sa12[p12]+n0], int(s[i0]), r12[i0/3])
	}
	return leq3(int(s[i12]), int(s[i12+1]), r12[sa12[p12]-n0+1], int(s[i0]), int(s[i0+1]), r12[i0/3+n0])
}

// `in`: input array containing the positions of the suffixes to be sorted
// `out`: output array where the sorted positions will be sorted
// `kth`: passed as 2, 1, then 0
// Essentially one pass of counting sort
func radixPass(s []byte, in, out []int, n02, kth, K int) {
	c := make([]int, K+1)

	// reset counts
	for i := 0; i <= K; i++ {
		c[i] = 0
	}

	// count occurrences of byte char
	for i := 0; i < n02; i++ {
		// `in[i]+kth` is suffix start+kth position
		// `s[^]` is the byte at that position in s
		// ^ byte is seen and the count incremented
		c[s[in[i]+kth]]++
	}

	// turn c into exclusive prefix sums
	for i, sum := 0, 0; i <= K; i++ {
		temp := c[i]
		c[i] = sum
		sum += temp // equals n02 at end
	}

	// store sorted positions
	for i := 0; i < n02; i++ {
		// `s[in[i]+kth]` is the kth byte at that suffix position
		// c[^] is the prefix sum of that byte
		prefix_sum := c[s[in[i]+kth]]
		out[prefix_sum] = in[i]
		// increment (part of counting sort)
		c[s[in[i]+kth]]++
	}
}

// lexicographic order for pairs
func leq2(a1, a2, b1, b2 int) bool {
	return (a1 < b1) || (a1 == b1 && a2 <= b2)
}

// lexicographic order for triples
func leq3(a1, a2, a3, b1, b2, b3 int) bool {
	return (a1 < b1) || (a1 == b1 && leq2(a2, a3, b2, b3))
}

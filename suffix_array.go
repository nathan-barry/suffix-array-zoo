package main

import (
	"fmt"
	"log"
)

// %%%%%%%%%% Initialization %%%%%%%%%%

const (
	NAIVE    = "naive"
	DC3_SKEW = "dc3_skew"
)

type SuffixArray struct {
	data []byte
	sa   []int
}

func NewSuffixArray(data []byte, algorithm string) SuffixArray {
	sa := SuffixArray{
		data: data,
		sa:   make([]int, len(data)),
	}

	switch algorithm {
	case NAIVE:
		sa.naiveSolution()
	case DC3_SKEW:
		sa.dc3Skew()
	default:
		log.Fatal("algorithm not implemented: not valid or typo in name")
	}

	return sa
}

// %%%%%%%%%% Search %%%%%%%%%%%%

func (sa *SuffixArray) Contains(text string) bool {
	l, r := 0, len(sa.data)-1

	for l <= r {
		m := (r + l) / 2
		suffix := (sa.data)[sa.sa[m] : sa.sa[m]+len(text)]

		if text == string(suffix) {
			return true
		} else if text < string(suffix) {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return false
}

func (sa *SuffixArray) Find(text string) int {
	l, r := 0, len(sa.data)-1

	for l <= r {
		m := (r + l) / 2
		suffix := sa.data[sa.sa[m] : sa.sa[m]+len(text)]

		if text == string(suffix) {
			return sa.sa[m]
		} else if text < string(suffix) {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return -1
}

// %%%%%%%%%% Printing %%%%%%%%%%

func (sa *SuffixArray) Print() {
	for _, s_i := range sa.sa { // s_i is the sorted suffix index
		fmt.Printf("%4d : %v\n", s_i, sa.data[s_i:])
	}
}

func (sa *SuffixArray) PrintTruncate(k int) {
	for _, s_i := range sa.sa {
		if s_i+k <= len(sa.data) {
			fmt.Printf("%4d : %v\n", s_i, sa.data[s_i:s_i+k])
		} else {
			fmt.Printf("%4d : %v\n", s_i, sa.data[s_i:])
		}
	}
}

package main

import (
	"fmt"
)

// %%%%%%%%%% Initialization %%%%%%%%%%

type SuffixArray struct {
	text    *string
	indices []int
}

func NewSuffixArray(text string) SuffixArray {
	SA := SuffixArray{
		text:    &text,
		indices: make([]int, len(text)),
	}
	return SA
}

// %%%%%%%%%% Search %%%%%%%%%%%%

func (sa *SuffixArray) Contains(text string) bool {
	l, r := 0, len(*sa.text)-1

	for l <= r {
		m := (r + l) / 2
		suffix := (*sa.text)[sa.indices[m] : sa.indices[m]+len(text)]

		if text == suffix {
			return true
		} else if text < suffix {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return false
}

func (sa *SuffixArray) Find(text string) int {
	l, r := 0, len(*sa.text)-1

	for l <= r {
		m := (r + l) / 2
		suffix := (*sa.text)[sa.indices[m] : sa.indices[m]+len(text)]

		if text == suffix {
			return sa.indices[m]
		} else if text < suffix {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return -1
}

// %%%%%%%%%% Printing %%%%%%%%%%

func (sa *SuffixArray) Print() {
	for _, i := range sa.indices {
		fmt.Printf("%4d : %v\n", i, (*sa.text)[i:])
	}
}

func (sa *SuffixArray) PrintTruncate(k int) {
	for _, i := range sa.indices {
		if i+k <= len(*sa.text) {
			fmt.Printf("%4d : %v\n", i, (*sa.text)[i:i+k])
		} else {
			fmt.Printf("%4d : %v\n", i, (*sa.text)[i:i])
		}
	}
}

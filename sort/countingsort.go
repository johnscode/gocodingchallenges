package sort

import (
	"slices"
)

func CountingSort(in []int) []int {
	if len(in) == 0 {
		return []int{}
	}

	// find the min/max values
	min := slices.Min(in)
	max := slices.Max(in)
	// create the count array
	counts := make([]int, max-min+1)
	for _, v := range in {
		counts[v-min]++
	}
	// determine cumulative counts
	for i := 1; i < len(counts); i++ {
		counts[i] = counts[i-1] + counts[i]
	}
	// create the output array
	out := make([]int, len(in))
	for i := range in {
		v := in[len(in)-1-i]
		out[counts[v-min]-1] = v
		counts[v-min]--
	}
	return out
}

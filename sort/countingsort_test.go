package sort

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	input    []int
	expected []int
}{
	{[]int{2, 5, 3, 0, 2, 3, 0, 3}, []int{0, 0, 2, 2, 3, 3, 3, 5}},
	{[]int{2, 5, 3, 1, 2, 3, 1, 3}, []int{1, 1, 2, 2, 3, 3, 3, 5}},
	{[]int{1, 1, 2, 4, 3, 0, 1, 2, 3, 1}, []int{0, 1, 1, 1, 1, 2, 2, 3, 3, 4}},
	{[]int{1, 3, 1, 2}, []int{1, 1, 2, 3}},
	{[]int{}, []int{}},
	{[]int{0}, []int{0}},
	{[]int{1}, []int{1}},
}

func TestCountingSort(t *testing.T) {

	for _, tc := range testCases {
		result := CountingSort(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("CountingSort(%v) = %v; expected %v", tc.input, result, tc.expected)
		}
	}
}

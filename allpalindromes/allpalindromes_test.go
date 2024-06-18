package allpalindromes

import (
	"reflect"
	"slices"
	"testing"
)

func TestFindAllPalindromes(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		// note that expected arrays have been presorted for quicker test runs
		{"", []string{}},
		{"a", []string{"a"}},
		{"ab", []string{"a", "b"}},
		{"aba", []string{"a", "aa", "aba", "b"}},
		{"aab", []string{"a", "aa", "b"}},
		{"abcba", []string{"a", "aa", "aba", "abba", "abcba", "aca", "b", "bb", "bcb", "c"}},
	}

	for _, tc := range testCases {
		results := FindAllPalindromes(tc.input)
		// sort result to match expected order
		slices.Sort(results)
		if !reflect.DeepEqual(results, tc.expected) {
			t.Errorf("findUniqueCombinations(%q) = %v; expected %v", tc.input, results, tc.expected)
		}
	}
}

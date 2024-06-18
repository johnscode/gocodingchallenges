package uniquecombos

import (
	"reflect"
	"testing"
)

func TestFindUniqueCombinations(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{"", []string{}},
		{"a", []string{"a"}},
		{"ab", []string{"a", "ab", "b"}},
		{"abc", []string{"a", "ab", "abc", "ac", "b", "bc", "c"}},
		{"aab", []string{"a", "aa", "aab", "ab", "b"}},
		{"aba", []string{"a", "ab", "aba", "aa", "b", "ba"}},
		{"abcd", []string{"a", "ab", "abc", "abcd", "abd", "ac", "acd", "ad", "b", "bc", "bcd", "bd", "c", "cd", "d"}},
	}

	for _, tc := range testCases {
		result := FindUniqueCombinations(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("findUniqueCombinations(%q) = %v; expected %v", tc.input, result, tc.expected)
		}
	}
}

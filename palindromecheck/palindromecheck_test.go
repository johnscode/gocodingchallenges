package palindromecheck

import "testing"

func TestPalindromeCheck(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"", true}, // one could argue this; but technically true
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"abc", false},
		{"aab", false},
		{"aba", true},
		{"a ba", true},
		{"ab a", true},
		{"a b a", true},
		{"abcd", false},
		{"ab cd", false},
		{"ab ba", true},
		{"ab b a", true},
		{"a b b a", true},
	}

	for _, tc := range testCases {
		if result := PalindromeCheck(tc.input); result != tc.expected {
			t.Errorf("PalindromeCheck: %s, expected %t", tc.input, tc.expected)
		}

	}
}

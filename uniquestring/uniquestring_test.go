package uniquestring

import "testing"

var testCases = []struct {
	input    string
	expected int
}{
	{"", 0},
	{"a", 1},
	{"abcdefg", 7},
	{"abracadabra", 5},
	{"Hello, World!", 10},
	{"111222333", 3},
	{"!@#$%^&()", 9},
	{"aabbccddee", 5},
	{"🍎🍌🍇🍓🍒", 5},
}

func TestCountUniqueChars(t *testing.T) {

	for _, tc := range testCases {
		result := CountUniqueChars(tc.input)
		if result != tc.expected {
			t.Errorf("countUniqueChars(%q) = %d; want %d", tc.input, result, tc.expected)
		}
	}
}

package maxprofit

import "testing"

func TestFindMaxProfit(t *testing.T) {
	testCases := []struct {
		input    []float64
		expected float64
	}{
		{[]float64{7, 1, 5, 3, 6, 4}, 7},
		{[]float64{7, 6, 4, 3, 1}, 0},
	}

	for _, testCase := range testCases {
		result := FindMaxProfit(testCase.input)
		if result != testCase.expected {
			t.Errorf("FindMaxProfit(%+v) = %f; expected %f", testCase.input, result, testCase.expected)
		}
	}
}

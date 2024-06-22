package bestbuysell1

import "testing"

func TestFindMaxProfit(t *testing.T) {
	testCases := []struct {
		input    []float64
		expected float64
	}{
		{[]float64{7, 1, 5, 3, 6, 4}, 5},
		{[]float64{7, 6, 4, 3, 1}, 0},
		{[]float64{8, 2, 6, 4, 7, 5, 1, 1}, 5},
		{[]float64{8, 2, 6, 4, 7, 5, 1, 6}, 5},
		{[]float64{8, 2, 6, 4, 7, 5, 1, 7}, 6},
	}

	for _, testCase := range testCases {
		result := FindSingleBestBuySell(testCase.input)
		if result != testCase.expected {
			t.Errorf("FindMaxProfit(%+v) = %f; expected %f", testCase.input, result, testCase.expected)
		}
	}
}

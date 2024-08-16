package gasstation

import "testing"

//var exampleData = [][]int{
//	{4, 6}, // gas at station, distance to next
//	{6, 5},
//	{7, 3},
//	{4, 5},
//}

func TestSolveGasStationProblem(t *testing.T) {
	testCases := []struct {
		input         [][]int
		expected      bool
		expectedStart int
	}{
		{
			[][]int{
				{4, 6},
				{6, 5},
				{7, 3},
				{4, 5},
			},
			true,
			1,
		},
		{
			[][]int{
				{1, 3},
				{2, 4},
				{3, 5},
				{4, 1},
				{5, 2},
			},
			true,
			3,
		},
	}
	for i, tc := range testCases {
		if result, start := SolveGasStationProblem(tc.input); result != tc.expected || start != tc.expectedStart {
			t.Errorf("SolveGasStationProblem: case %d, expected %t and start %d", i, tc.expected, tc.expectedStart)
		}

	}

}

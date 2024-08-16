package gasstation

/**
Problem Statement:
Imagine a circular route with N gas stations. Each gas station has a certain amount of gas and a distance to the next station. You have a vehicle that can store an unlimited amount of gas. The goal is to find the starting gas station (if it exists) from which you can complete a full circular tour, visiting all stations and returning to the starting point without running out of gas.

Key Points:
1. The route is circular, meaning the last station connects back to the first station.
2. At each station, you can fill your tank with the available gas.
3. You consume gas while traveling between stations.
4. You need to find a starting point that allows you to complete the entire circuit.

Constraints:
- You can only move in one direction (clockwise).
- The amount of gas at each station and the distance to the next station are given.
- If no solution exists, you should be able to determine that as well.

Example:
Consider 4 gas stations with the following data:
Station 1: Gas = 4, Distance to next = 6
Station 2: Gas = 6, Distance to next = 5
Station 3: Gas = 7, Distance to next = 3
Station 4: Gas = 4, Distance to next = 5
*/

/*
	exampleData:
		0: gas available at station
		1: distance to next

var exampleData = [][]int{
	{4,6},
	{6,5},
	{7,3},
	{4,5},
}

*/

func SolveGasStationProblem(stations [][]int) (bool, int) {
	if !enoughTotalGas(stations) {
		return false, 0
	}

	tank := 0
	surplus := 0
	start := 0
	for i, station := range stations {
		surplus += station[0] - station[1]
		tank += station[0] - station[1]
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}
	if surplus >= 0 {
		return true, start
	}
	return false, 0
}

func enoughTotalGas(stations [][]int) bool {
	surplus := 0
	for _, station := range stations {
		surplus += station[0] - station[1]
	}
	return surplus >= 0
}

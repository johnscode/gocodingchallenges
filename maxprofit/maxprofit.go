package maxprofit

func FindMaxProfit(prices []float64) float64 {
	maxProfit := 0.0
	// if you only have a single price, then profit is 0
	if len(prices) <= 1 {
		return 0
	}
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}

package bestbuysell1

func FindSingleBestBuySell(prices []float64) float64 {
	buy := prices[0]
	maxProfit := 0.0
	for i := 1; i < len(prices); i++ {
		if buy > prices[i] {
			buy = prices[i]
		} else if prices[i]-buy > maxProfit {
			maxProfit = prices[i] - buy
		}
	}
	return maxProfit
}

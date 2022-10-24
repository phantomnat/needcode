package easy

func (p *Practice) MaximizeProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	maxPrices := make([]int, len(prices))
	maxPrices[len(prices)-1] = 0

	for i := len(prices) - 2; i >= 0; i-- {
		maxPrices[i] = prices[i+1]
		if maxPrices[i+1] > maxPrices[i] {
			maxPrices[i] = maxPrices[i+1]
		}
	}
	// prices : 7 1 5 3 6 4
	// max    :   6 6 6 4 0
	profit := 0
	for i := range prices {
		p := maxPrices[i] - prices[i]
		if p  > profit {
			profit = p
		}
	}
	return profit
}

package algorithm

func maxProfit(prices []int) int {
	answer := 0
	// 每天可以买卖无数次，则只需要当天比前一天买入的价格高，就可以在前一天买入，在当天卖出即可
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			answer += prices[i] - prices[i-1]
		}
	}
	return answer
}

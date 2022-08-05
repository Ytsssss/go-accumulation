package algorithm

func CoinChange(coins []int, amount int) int {
	// dp[i]表示到达整数i所需要的最少硬币数
	// dp[i]=min{dp[i-coins[j]]}+1
	dp := make([]int, amount+1)
	maxCount := amount + 1
	for i := 1; i < len(dp); i++ {
		dp[i] = maxCount
	}
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

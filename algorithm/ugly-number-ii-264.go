package algorithm

func NthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		cur2, cur3, cur5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(min(cur2, cur3), cur5)
		if dp[i] == cur2 {
			p2++
		}
		if dp[i] == cur3 {
			p3++
		}
		if dp[i] == cur5 {
			p5++
		}
	}
	return dp[n]
}

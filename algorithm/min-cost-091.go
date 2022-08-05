package algorithm

func MinCost(costs [][]int) int {
	answer := make([][]int, len(costs))
	answer[0] = costs[0]
	for i := 1; i < len(costs); i++ {
		tmp := make([]int, 3)
		for k := 0; k <= 2; k++ {
			tmp[k] = min(answer[i-1][(k+1)%3], answer[i-1][(k+2)%3]) + costs[i][k]
		}
		answer[i] = tmp
	}

	return min(min(answer[len(costs)-1][0], answer[len(costs)-1][1]), min(answer[len(costs)-1][2], answer[len(costs)-1][1]))
}

func min(num1, num2 int) int {
	if num1 > num2 {
		return num2
	}
	return num1
}

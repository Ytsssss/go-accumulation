package algorithm

func LenLongestFibSubseq(arr []int) int {
	// 存储数组元素和下标关系，由于数组严格递增，则该map的key设置为数组值
	numIndexMap := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		numIndexMap[arr[i]] = i
	}
	// 定义dp状态，dp[i][j]表示选取i，j为斐波拉契最后两位
	// 所以nums[j]-nums[i]=x,x为数组中的一个数，如果存在该数index为k，则 dp[i][j]=dp[k][i]+1
	// 最后遍历dp数组，获取最大值
	dp := make([][]int, len(arr))
	maxCount := 0
	for i := 0; i < len(arr); i++ {
		arri := make([]int, len(arr))
		dp[i] = arri
		for j := i + 1; j < len(arr); j++ {
			count := arr[j] - arr[i]
			if k, ok := numIndexMap[count]; ok && k < i {
				// 此处必须保证k的值小于i，表示k在i之前；此处和3进行比较是因为一个斐波拉契序列最小长度为3
				dp[i][j] = max(dp[k][i]+1, 3)
				maxCount = max(dp[i][j], maxCount)
			}
		}
	}
	return maxCount
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

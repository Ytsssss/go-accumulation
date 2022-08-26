package algorithm

func SubarraySum(nums []int, k int) int {
	count := 0
	left, right := 0, 0
	sum := nums[0]
	for right < len(nums) {
		for sum < k && right < len(nums)-1 {
			right++
			sum += nums[right]
		}
		if sum == k {
			count++
			sum -= nums[left]
			left++
			sum += nums[right]
			right++
			continue
		}
		for sum > k {
			sum -= nums[left]
			left++
		}
	}
	return count
}

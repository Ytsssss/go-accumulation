package algorithm

func FindKthLargest(nums []int, k int) int {
	index := len(nums) - k
	return find(nums, 0, len(nums)-1, index)
}

func find(nums []int, left, right, k int) int {
	ans := -1
	p := partition0(nums, left, right)
	if p == k {
		return nums[k]
	} else if p > k {
		ans = find(nums, left, p-1, k)
	} else {
		ans = find(nums, p+1, right, k)
	}
	return ans
}

func partition0(nums []int, left, right int) int {
	num, tmp := nums[left], left
	for left < right {
		for nums[right] >= num && left < right {
			right--
		}
		for nums[left] <= num && left < right {
			left++
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	nums[left], nums[tmp] = nums[tmp], nums[left]
	return left
}

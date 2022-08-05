package algorithm

import "sort"

func FindPairs(nums []int, k int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)
	answer := 0
	i, j := 0, 1
	lasti := nums[i] - 1
	for j < len(nums) {
		if nums[j]-nums[i] == k {
			if nums[i] != lasti {
				answer++
				lasti = nums[i]
			}
			i++
			j++
		} else if nums[j]-nums[i] > k {
			i++
		} else {
			j++
		}
		if i == j {
			j++
		}
	}

	return answer
}

package algorithm

import "sort"

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}
	sort.Ints(nums)
	var answer [][]int
	for i := 0; i < len(nums); i++ {
		// 此处判断是否和前一个数相同，是为了避免出现重复结果
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, len(nums)-1
			for left < right {
				if left > j+1 && nums[left] == nums[left-1] {
					left++
					continue
				}
				if nums[i]+nums[j]+nums[left]+nums[right] == target {
					answer = append(answer, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					right--
				} else if nums[i]+nums[j]+nums[left]+nums[right] < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return answer
}

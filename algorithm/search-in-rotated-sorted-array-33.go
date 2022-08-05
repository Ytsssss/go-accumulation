package algorithm

func Search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	i, j := 0, len(nums)-1
	for i <= j {
		middle := i + (j-i)/2
		if nums[middle] == target {
			return middle
		}
		if nums[0] <= nums[middle] {
			if target < nums[middle] && nums[0] <= target {
				j = middle - 1
			} else {
				i = middle + 1
			}
		} else {
			if target > nums[middle] && target <= nums[len(nums)-1] {
				i = middle + 1
			} else {
				j = middle - 1
			}
		}
	}

	return -1
}

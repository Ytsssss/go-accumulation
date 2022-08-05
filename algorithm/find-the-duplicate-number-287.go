package algorithm

func FindDuplicate(nums []int) int {
	i := 1
	for i <= len(nums) {
		for nums[i-1] != i {
			if nums[i-1] == nums[nums[i-1]-1] {
				return nums[i-1]
			}
			nums[i-1], nums[nums[i-1]-1] = nums[nums[i-1]-1], nums[i-1]
		}
		i++
	}
	return -1
}

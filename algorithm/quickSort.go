package algorithm

func SortArray(nums []int) []int {
	left, right := 0, len(nums)-1
	quickSort(nums, left, right)
	return nums
}

func partition(nums []int, left, right int) int {
	pilot, tmp := nums[left], left
	for left < right {
		for left < right && pilot <= nums[right] {
			right--
		}
		for left < right && pilot >= nums[left] {
			left++
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	nums[left], nums[tmp] = nums[tmp], nums[left]
	return left
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	pilot := partition(nums, left, right)
	quickSort(nums, left, pilot-1)
	quickSort(nums, pilot+1, right)
}

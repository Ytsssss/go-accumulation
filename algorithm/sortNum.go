package algorithm

func getMiddle(nums []int) float32 {
	if len(nums) < 1 {
		return -1
	}
	sortNums(nums)
	if len(nums)%2 == 0 {
		return float32((nums[len(nums)/2] + nums[len(nums)/2+1]) / 2)
	}
	return float32(nums[len(nums)/2])
}

func sortNums(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

func GetMiddle2(nums []int) float32 {
	if len(nums) < 1 {
		return -1
	}
	index := len(nums) / 2
	left, right := 0, len(nums)-1
	for left < index && right > index {
		if nums[left] < nums[index] {
			left++
		} else {
			nums[left], nums[index] = nums[index], nums[left]
		}
		if nums[right] > nums[index] {
			right--
		} else {
			nums[right], nums[index] = nums[index], nums[right]
		}
	}
	return float32(nums[index])
}

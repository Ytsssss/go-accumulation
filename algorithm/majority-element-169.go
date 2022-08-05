package algorithm

func majorityElement(nums []int) int {
	countMap := map[int]int{}
	for _, num := range nums {
		if i, _ := countMap[num]; i+1 >= (len(nums)+1)/2 {
			return num
		} else {
			countMap[num] = i + 1
		}
	}
	return 0
}

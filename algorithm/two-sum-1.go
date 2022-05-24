package algorithm

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)
	for i, num := range nums {
		if j, ok := m[target-num]; ok && i != j {
			return []int{i, j}
		}
		m[num] = i
	}
	return nil
}

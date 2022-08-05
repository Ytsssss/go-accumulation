package algorithm

func IsHappy(n int) bool {
	mapK := make(map[int]bool, 0)

	for n > 0 {
		if n == 1 {
			return true
		}
		if mapK[n] {
			return false
		}

		mapK[n] = true
		n = getNew(n)
	}
	return n == 1
}

func getNew(n int) int {
	result := 0
	i := 0
	k := n
	for k > 0 {
		i = k % 10
		k = k / 10
		result += i * i
	}
	return result
}

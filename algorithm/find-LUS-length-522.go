package algorithm

func FindLUSlength(strs []string) int {
	max := -1
	for i := 0; i < len(strs); i++ {
		check := false
		for j := 0; j < len(strs); j++ {
			if i != j && checkSuf(strs[i], strs[j]) {
				check = true
				break
			}
		}
		if !check && len(strs[i]) > max {
			max = len(strs[i])
		}
	}
	return max
}

func checkSuf(str1, str2 string) bool {
	i, j := 0, 0
	for i < len(str1) && j < len(str2) {
		if str1[i] == str2[j] {
			i++
			j++
		} else {
			if len(str1)-i < len(str2)-j {
				j++
			} else {
				break
			}
		}
	}
	return i == len(str1)
}

func getLength(a, b string) int {
	if len(a) > len(b) {
		return len(a)
	}
	return len(b)
}

func findLUSlength1(a string, b string) int {
	if len(a) > len(b) {
		return len(a)
	} else if len(a) < len(b) {
		return len(b)
	}
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
			j++
		} else {
			break
		}
	}
	if i == len(a) && j == len(b) {
		return -1
	}

	return len(a)
}

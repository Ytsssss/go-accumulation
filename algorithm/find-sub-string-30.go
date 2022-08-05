package algorithm

func FindSubstring(s string, words []string) []int {
	length := 0
	n := len(words[0])
	valueMap := make(map[string]int, 0)
	for _, word := range words {
		length += len(word)
		valueMap[word] = valueMap[word] + 1
	}
	result := make([]int, 0)
	i, j := 0, length
	for j <= len(s) {
		tmp := s[i:j]
		tmpMap := make(map[string]int)
		for key, value := range valueMap {
			tmpMap[key] = value
		}
		for k := 0; k <= len(tmp)-n; k = k + n {
			if num, ok := tmpMap[tmp[k:k+n]]; ok && num > 0 {
				if num == 1 {
					delete(tmpMap, tmp[k:k+n])
				} else {
					tmpMap[tmp[k:k+n]] = num - 1
				}
			} else {
				break
			}
		}
		if len(tmpMap) == 0 {
			result = append(result, i)
		}
		i = i + 1
		j = j + 1
	}
	return result
}

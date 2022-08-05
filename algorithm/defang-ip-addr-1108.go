package algorithm

func defangIPaddr(address string) string {
	var answer string
	for _, item := range address {
		if item != '.' {
			answer += string(item)
		} else {
			answer += "[.]"
		}
	}
	return answer
}

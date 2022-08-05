package algorithm

func LengthOfLongestSubstring(s string) int {
	left, right, charMap, maxCount := 0, 0, make(map[byte]bool), 0
	for right < len(s) {
		rightChar := s[right]
		if charMap[rightChar] {
			for {
				if s[left] != rightChar {
					delete(charMap, s[left])
					left++
				} else {
					left++
					break
				}
			}
		} else {
			charMap[rightChar] = true
		}
		right++
		maxCount = max(maxCount, right-left)
	}
	return maxCount
}

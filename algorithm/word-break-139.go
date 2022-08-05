package algorithm

func WordBreak(s string, wordDict []string) bool {
	length := len(s)
	dp := make([]bool, length+1)
	wordMap := make(map[string]bool)
	for i := 0; i < len(wordDict); i++ {
		wordMap[wordDict[i]] = true
	}
	dp[0] = true
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= length; j++ {
			tmp := s[i:j]
			if wordMap[tmp] {
				dp[j] = dp[i]
			}
		}
	}
	return dp[length]
}

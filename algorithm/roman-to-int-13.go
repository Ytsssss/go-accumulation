package algorithm

func RomanToInt(s string) int {
	strMap := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	specialMap := map[byte]byte{
		'V': 'I',
		'X': 'I',
		'L': 'X',
		'C': 'X',
		'D': 'C',
		'M': 'C',
	}
	answer := 0
	for i := len(s) - 1; i >= 0; i-- {
		if v, ok := specialMap[s[i]]; ok {
			if i > 0 && s[i-1] == v {
				answer += strMap[s[i-1:i+1]]
				i--
				continue
			}
		}
		answer += strMap[string(s[i])]
	}
	return answer
}

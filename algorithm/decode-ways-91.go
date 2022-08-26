package algorithm

import (
	"fmt"
	"strconv"
)

func NumDecodings(s string) int {
	answer := 1
	if s[0] == '0' {
		return 0
	}
	for i := 0; i < len(s)-1; i++ {
		if s[i+1] == '0' {
			if s[i]-'2' > 0 {
				return 0
			} else {
				i++
				continue
			}
		}
		if s[i] == '1' {
			answer++
		}
		if s[i] == '2' && s[i+1] <= '6' {
			answer++
		}
	}
	atoi, _ := strconv.Atoi("001")
	fmt.Println(atoi)
	return answer
}

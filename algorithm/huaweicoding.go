package algorithm

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		fmt.Print(getResultHJ1(input.Text()))
	}
}

func getResultHJ1(input string) int {
	nums := strings.Split(input, " ")
	return len(nums[len(nums)-1])
}

func CountStr(str string, ch string) int {
	str = strings.ToLower(str)
	charMap := make(map[byte]int)
	for i := 0; i < len(str); i++ {
		charMap[str[i]]++
	}
	ch = strings.ToLower(ch)
	return charMap[ch[0]]
}

func GetRandomCount() {
	var n int
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		n, _ = strconv.Atoi(scanner.Text())
	}
	nums := make([]int, 501)
	for scanner.Scan() {
		k, _ := strconv.Atoi(scanner.Text())
		nums[k] = 1
		n--
		if n == 0 {
			break
		}
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] != 0 {
			fmt.Println(i)
		}
	}
}

func PrintStr() {
	var str string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		str = scanner.Text()
	}
	i := 0
	for i < len(str)-8 {
		fmt.Println(str[i : i+8])
		i = i + 8
	}
	if i < len(str) {
		tmp := str[i:]
		for k := 0; k < 8-(len(str)-i); k++ {
			tmp += "0"
		}
		fmt.Println(tmp)
	}
}

func DecimbalChange() {
	var str string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		str = scanner.Text()
	}
	parseInt, _ := strconv.ParseInt(str, 0, 32)
	fmt.Println(parseInt)
}

func QualityFactor() {
	var num int
	fmt.Scanln(&num)
	for i := 2; i*i <= num; i++ {
		for num%i == 0 {
			num = num / i
			fmt.Print(num)
			fmt.Print(" ")
		}
	}
	if num != 1 {
		fmt.Print(num)
	}
}

func Similar() {
	var num float64
	fmt.Scanln(&num)
	str := strconv.FormatFloat(num, 'f', 32, 32)
	split := strings.Split(str, ".")
	i, _ := strconv.Atoi(split[0])
	atoi, _ := strconv.Atoi(string(split[1][0]))
	if atoi >= 5 {
		i = i + 1
	}
	fmt.Println(i)
}

func MergeRecord() {
	var n int
	var k, v int
	fmt.Scanln(&n)
	countMap := make(map[int]int)
	for n > 0 {
		fmt.Scanf("%d %d", &k, &v)
		countMap[k] += v
		n--
	}
	keys := make([]int, 0, len(countMap))
	for key := range countMap {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for key := range keys {
		fmt.Print(keys[key])
		fmt.Print(" ")
		fmt.Println(countMap[keys[key]])
	}
}

func getUnRepeat() {
	var num int
	fmt.Scanln(&num)
	numMap := make(map[int]bool)
	result := 0
	for num > 0 {
		tmp := num % 10
		if !numMap[tmp] {
			result = result*10 + tmp
		}
		numMap[tmp] = true
		num = num / 10
	}
	fmt.Println(result)
}

func reverseNum() {
	var num int
	fmt.Scanln(&num)
	if num == 0 {
		fmt.Println(num)
		return
	}
	for num != 0 {
		fmt.Print(num % 10)
		num = num / 10
	}
}

func reverseStr() {
	var str string
	fmt.Scanln(&str)
	index := 0
	for index < len(str) {
		fmt.Print(string(str[len(str)-index-1]))
		index++
	}
}

func reverseDoc() {
	var str string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		str = scanner.Text()
	}
	split := strings.Split(str, " ")
	index := 0
	for index < len(split) {
		fmt.Print(split[len(split)-index-1])
		fmt.Print(" ")
		index++
	}
}

func sortStr() {
	var num int
	var str string
	fmt.Scanln(&num)
	strs := make([]string, 0, num)
	for num > 0 {
		fmt.Scanln(&str)
		strs = append(strs, str)
		num--
	}
	sort.Strings(strs)
	for i := 0; i < len(strs); i++ {
		fmt.Println(strs[i])
	}
}

func countInt() {
	var num int
	fmt.Scanln(&num)
	answer := 0
	for num > 0 {
		num = num & (num - 1)
		answer++
	}
	fmt.Println(answer)
}

func Check() {
	var str string
	fmt.Scanln(&str)
	first, _ := strconv.ParseInt(string(str[0]), 16, 32)
	second, _ := strconv.ParseInt(string(str[1]), 16, 32)
	num := first<<4 + second
	if len(str) == 2 {
		if num&127 != num {
			fmt.Println(-1)
			return
		}
		fmt.Println(num & 127)
		return
	}

	if len(str) == 4 {
		if num&223 != num {
			fmt.Println(-1)
			return
		}
		num = num & 31
		i := getNum(str, 2)
		if i == -1 {
			fmt.Println(-1)
			return
		}
		fmt.Println(num<<6 + i)
		return
	}
	if len(str) == 6 {
		if num&239 != num {
			fmt.Println(-1)
			return
		}
		num = num & 15
		index := 2
		for index <= 4 {
			i := getNum(str, index)
			if i == -1 {
				fmt.Println(-1)
				return
			}
			num = num<<6 + i
			index += 2
		}
		fmt.Println(num)
	}
	if len(str) == 8 {
		if num&247 != num {
			fmt.Println(-1)
			return
		}
		num = num & 7
		index := 2
		for index <= 6 {
			i := getNum(str, index)
			if i == -1 {
				fmt.Println(-1)
				return
			}
			num = num<<6 + i
			index += 2
		}
		fmt.Println(num)
	}
	if len(str) == 10 {
		if num&251 != num {
			fmt.Println(-1)
			return
		}
		num = num & 3
		index := 2
		for index <= 8 {
			i := getNum(str, index)
			if i == -1 {
				fmt.Println(-1)
				return
			}
			num = num<<6 + i
			index += 2
		}
		fmt.Println(num)
	}
	if len(str) == 10 {
		if num&253 != num {
			fmt.Println(-1)
			return
		}
		index := 2
		num = num & 1
		for index <= 10 {
			i := getNum(str, index)
			if i == -1 {
				fmt.Println(-1)
				return
			}
			num = num<<6 + i
			index += 2
		}
		fmt.Println(num)
	}
}

func getNum(str string, i int) int64 {
	first, _ := strconv.ParseInt(string(str[i]), 16, 32)
	second, _ := strconv.ParseInt(string(str[i+1]), 16, 32)
	num := first<<4 + second
	if num&191 != num {
		return -1
	}
	return num & 63
}

func getAnswert() {
	var num int
	fmt.Scanln(&num)
	var count int
	fmt.Scanln(&count)
	for count > 0 {
		var a, b int
		fmt.Scanf("%d>%d", a, b)

	}
}

type node struct {
	nums [][]int
}

func (n *node) union(a, b int) {
	n.nums[a][b] = 1
	n.nums[b][a] = 1

}

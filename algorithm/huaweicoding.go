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

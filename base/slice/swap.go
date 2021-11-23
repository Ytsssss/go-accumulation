package slice

import "fmt"

func swapError1() {
	a := []int{1, 2, 3, 4}
	defer func(a []int) {
		fmt.Printf("a: %v\n", a)
	}(a)
	a[0], a[1] = a[2], a[4]
	// a: [1 2 3 4]
	// index out of range [4] with length 4
	// 第一步计算等号左侧索引表达式和取址表达式，计算等号右侧的表达式的值(panic)
	// 第二步赋值操作没有进行，所以a还是原值
}

func swapError2() {
	a := []int{1, 2, 3, 4}
	defer func(a []int) {
		fmt.Printf("a: %v\n", a)
	}(a)
	a[0], a[4] = a[2], a[3]
	// a: [3 2 3 4]
	// index out of range [4] with length 4
	// 第一步会将表达式转化为 a[0],a[4]=3,4
	// 第二步进行赋值操作，先对a[0]进行赋值，再对a[4]进行赋值(panic)，
	// 所以结果a=[]int{3,2,3,4}
}

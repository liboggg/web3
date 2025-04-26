package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
}

func singleNumber(nums []int) int {
	//1. 使用map 计算数量
	// var m = make(map[int]int)
	// for _, v := range nums {
	// 	value := m[v]
	// 	if value == 1 {
	// 		delete(m, v)
	// 	} else {
	// 		m[v] = 1
	// 	}
	// }
	// fmt.Println(m)
	// for k, _ := range m {
	// 	return k
	// }
	// return 0

	//2. 使用异或
	//相同数异或=0
	//0与任何数异或=任意数
	//异或满足交换律，所以异或所有数，最后剩下的数就是只出现一次的数
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

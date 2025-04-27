package main

import "fmt"

// ✅指针
// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。
func main() {
	mulit1([]int{1, 2, 3, 4, 5})
	fmt.Println(mulit2([]int{1, 2, 3, 4, 5}))
	fmt.Println(mulit3(&[]int{1, 2, 3, 4, 5}))
}

func mulit1(nums []int) {
	for _, value := range nums {
		fmt.Println(value * 2)
	}
}

func mulit2(nums []int) []int {
	fmt.Println("mulit2", len(nums))
	newArr := make([]int, 0, len(nums))
	for _, value := range nums {
		newArr = append(newArr, value*2)
	}
	return newArr
}

func mulit3(nums *[]int) []int {
	fmt.Println("mulit3", len(*nums))
	newArr := make([]int, 0, len(*nums))
	for _, value := range *nums {
		newArr = append(newArr, value*2)
	}
	return newArr
}

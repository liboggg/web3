package main

import "fmt"

// ✅指针
// 题目 ：编写一个Go程序，定义一个函数，
// 该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，
// 然后在主函数中调用该函数并输出修改后的值。
func main() {
	num := 10
	add(&num)
	fmt.Println(num)
}

func add(a *int) {
	*a += 10
}

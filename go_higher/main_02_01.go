package main

import "time"

// ✅Goroutine
// 题目 ：编写一个程序，使用 go 关键字启动两个协程，
// 一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func main() {
	go print1()
	go print2()
	time.Sleep(10000 * time.Second)
}

func print1() {
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			go println("print1", i)
		}
	}
}

func print2() {
	for i := 2; i <= 10; i++ {
		if i%2 == 0 {
			go println("print2", i)
		}
	}
}

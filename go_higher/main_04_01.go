package main

import (
	"fmt"
	"time"
)

// ✅Channel
// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。
// 一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。
func main() {
	go producer()
	go consumer()
	time.Sleep(10000 * time.Second)
}

var queue = make(chan int)

func producer() {
	for i := 1; i <= 10; i++ {
		queue <- i
	}
}

func consumer() {
	for {
		pop := <-queue
		fmt.Println(pop)
	}
}

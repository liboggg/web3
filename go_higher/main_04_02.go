package main

import (
	"fmt"
	"time"
)

// ✅Channel
// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。
func main() {
	go producer2()
	go consumer2()
	time.Sleep(10000 * time.Second)
}

var queue2 = make(chan int, 2)

func producer2() {
	for i := 1; i <= 100; i++ {
		queue2 <- i
	}
}

func consumer2() {
	for {
		pop := <-queue2
		fmt.Println(pop)
	}
}

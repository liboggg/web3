package main

import "fmt"

func main() {
	fmt.Println(rob([]int{2, 1, 1, 2}))
}

func rob(nums []int) int {
	var a, b int

	for k, v := range nums {
		if k%2 == 0 {
			a += v
		} else {
			b += v
		}
	}

	return max(a, b)
}

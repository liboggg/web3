package main

import "fmt"

// 46. 全排列：给定一个不含重复数字的数组 nums ，返回其所有可能的全排列。
// 可以使用回溯算法，定义一个函数来进行递归操作，在函数中通过交换数组元素的位置来生成不同的排列
// ，使用 for 循环遍历数组，每次选择一个元素作为当前排列的第一个元素，然后递归调用函数处理剩余的元素。
func main() {
	fmt.Println(permute([]int{2, 2, 1}))
}
func permute(nums []int) [][]int {
	var res [][]int
	backtrack(&res, nums, 0)
	return res
}
func backtrack(res *[][]int, nums []int, index int) {
	if index == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}
	for i := index; i < len(nums); i++ {
		nums[i], nums[index] = nums[index], nums[i]
		backtrack(res, nums, index+1)
		nums[i], nums[index] = nums[index], nums[i]
	}

}

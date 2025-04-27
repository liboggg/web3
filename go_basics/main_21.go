package main

import "fmt"

// 21. 合并两个有序链表：将两个升序链表合并为一个新的升序链表并返回。
// 新链表是通过拼接给定的两个链表的所有节点组成的。可以定义一个函数，接收两个链表的头节点作为参数，在函数内部使用双指针法，
// 通过比较两个链表节点的值，将较小值的节点添加到新链表中，直到其中一个链表为空，然后将另一个链表剩余的节点添加到新链表中。
func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}

package main

import "fmt"

// 430. 扁平化多级双向链表：多级双向链表中，除了指向下一个节点和前一个节点指针之外，
// 它还有一个子链表指针，可能指向单独的双向链表。这些子列表也可能会有一个或多个自己的子项，
// 依此类推，生成多级数据结构，如下面的示例所示。给定位于列表第一级的头节点，请扁平化列表，
// 即将这样的多级双向链表展平成普通的双向链表，使所有结点出现在单级双链表中。
// 可以定义一个结构体来表示链表节点，包含 val、prev、next 和 child 指针，然后使用递归的方法来扁平化链表，
// 先处理当前节点的子链表，再将子链表插入到当前节点和下一个节点之间。
func main() {
	// 创建一个多级双向链表
	root := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node5 := &Node{Val: 5}
	node6 := &Node{Val: 6}
	node7 := &Node{Val: 7}
	node8 := &Node{Val: 8}
	node9 := &Node{Val: 9}
	node10 := &Node{Val: 10}
	node11 := &Node{Val: 11}
	node12 := &Node{Val: 12}

	root.Next = node2
	node2.Prev = root
	node2.Next = node3
	node3.Prev = node2
	node3.Next = node4
	node4.Prev = node3
	node4.Next = node5
	node5.Prev = node4
	node5.Next = node6
	node6.Prev = node5

	node3.Child = node7
	node7.Next = node8
	node8.Prev = node7
	node8.Next = node9
	node9.Prev = node8
	node9.Next = node10
	node10.Prev = node9

	node8.Child = node11
	node11.Next = node12
	node12.Prev = node11

	// 扁平化链表
	flattened := flatten(root)

	// 打印扁平化后的链表
	printList(flattened)
}

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	flattenHelper(root)
	return root

}

// flattenHelper 辅助函数，递归地扁平化链表
func flattenHelper(node *Node) *Node {
	if node == nil {
		return nil
	}

	current := node
	for current != nil {
		if current.Child != nil {
			// 保存 next 节点
			next := current.Next

			// 扁平化子链表
			flattenedChild := flattenHelper(current.Child)
			current.Child = nil

			// 将扁平化后的子链表插入到当前节点和 next 节点之间
			current.Next = flattenedChild
			flattenedChild.Prev = current

			// 找到扁平化子链表的最后一个节点
			lastChild := flattenedChild
			for lastChild.Next != nil {
				lastChild = lastChild.Next
			}

			// 连接最后一个子节点和 next 节点
			if next != nil {
				lastChild.Next = next
				next.Prev = lastChild
			}

			// 移动到下一个节点
			current = next
		} else {
			current = current.Next
		}
	}
	return node
}

// 打印链表，用于测试
func printList(head *Node) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

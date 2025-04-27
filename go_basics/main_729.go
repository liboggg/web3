package main

import "fmt"

// 729. 我的日程安排表 I：实现一个 MyCalendar 类来存放你的日程安排。
// 如果要添加的日程安排不会造成 重复预订 ，则可以存储这个新的日程安排。
// 当两个日程安排有一些时间上的交叉时（例如两个日程安排都在同一时间内），就会产生 重复预订 。
// 日程可以用一对整数 start 和 end 表示，这里的时间是半开区间，即 [start, end) ，
// 实数 x 的范围为 start <= x < end 。实现 MyCalendar 类：MyCalendar() 初始化日历对象。
// boolean book(int start, int end) 如果可以将日程安排成功添加到日历中而不会导致重复预订，
// 返回 true ，否则，返回 false 并且不要将该日程安排添加到日历中。可以定义一个结构体来表示日程安排，
// 包含 start 和 end 字段，然后使用一个切片来存储所有的日程安排，在 book 方法中，
// 遍历切片中的日程安排，判断是否与要添加的日程安排有重叠。
func main() {
	// 测试 MyCalendar 类
	obj := Constructor()
	fmt.Println(obj.Book(10, 20)) // true
	fmt.Println(obj.Book(15, 25)) // false
	fmt.Println(obj.Book(20, 30)) // true
}

// MyCalendar 结构体，包含一个切片来存储所有的日程安排
type MyCalendar struct {
	events [][]int
}

// Constructor 初始化 MyCalendar 对象
func Constructor() MyCalendar {
	return MyCalendar{
		events: [][]int{},
	}
}

// Book 方法检查新日程是否与现有日程重叠，并在不重叠的情况下添加新日程
func (this *MyCalendar) Book(startTime int, endTime int) bool {
	for _, event := range this.events {
		// 检查是否有重叠
		if startTime < event[1] && endTime > event[0] {
			return false
		}
	}
	// 没有重叠，添加新日程
	this.events = append(this.events, []int{startTime, endTime})
	return true
}

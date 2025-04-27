package main

import "fmt"

// ✅面向对象
// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
// 然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
// 在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。
func main() {
	a := Rectangle{}
	b := Circle{}
	a.Area()
	a.Perimeter()
	b.Area()
	b.Perimeter()
}

type Rectangle struct {
}

func (a Rectangle) Area() {
	fmt.Println("Rectangle Area")
}
func (a Rectangle) Perimeter() {
	fmt.Println("Rectangle Perimeter")
}

type Circle struct {
}

func (a Circle) Area() {
	fmt.Println("Circle Area")
}
func (a Circle) Perimeter() {
	fmt.Println("Circle Perimeter")
}

type Shape interface {
	Area()
	Perimeter()
}

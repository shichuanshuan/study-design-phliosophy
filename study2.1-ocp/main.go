package main

import (
	"fmt"
	"math"
)

// Shape 接口：所有图形都必须实现该接口
type Shape interface {
	Area() float64
}

// Rectangle 结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// 实现 Shape 接口
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle 结构体
type Circle struct {
	Radius float64
}

// 实现 Shape 接口
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 新增：Triangle（新需求，只扩展，不修改原来的代码）
type Triangle struct {
	Base   float64
	Height float64
}

// 实现 Shape 接口
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// AreaCalculator 处理一组图形
type AreaCalculator struct{}

// ComputeAreas 计算所有图形面积
func (ac AreaCalculator) ComputeAreas(shapes []Shape) {
	for _, shape := range shapes {
		fmt.Printf("Area: %.2f\n", shape.Area())
	}
}

/*
说明:
Shape 是抽象接口，对扩展开放。

Rectangle、Circle 和 Triangle 都是对接口的实现，属于扩展。

AreaCalculator 并不知道具体的图形种类，只依赖 Shape 接口。

后续增加新图形，只需新增结构体实现 Shape 接口即可，无需修改已有逻辑。
*/
func main() {
	shapes := []Shape{
		Rectangle{Width: 10, Height: 5},
		Circle{Radius: 7},
		Triangle{Base: 6, Height: 4}, // 添加新类型，不需修改 AreaCalculator
	}

	calculator := AreaCalculator{}
	calculator.ComputeAreas(shapes)
}

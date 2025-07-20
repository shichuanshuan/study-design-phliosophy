package main

import "fmt"

// 抽象接口
type Bird interface {
	Fly()
}

// 子类1：麻雀
type Sparrow struct{}

func (s Sparrow) Fly() {
	fmt.Println("Sparrow is flying")
}

// 子类2：燕子
type Swallow struct{}

func (s Swallow) Fly() {
	fmt.Println("Swallow is flying")
}

// 飞行测试函数，接受任何 Bird 类型
func LetBirdFly(b Bird) {
	b.Fly()
}

func main() {
	sparrow := Sparrow{}
	swallow := Swallow{}

	LetBirdFly(sparrow) // 输出：Sparrow is flying
	LetBirdFly(swallow) // 输出：Swallow is flying
}

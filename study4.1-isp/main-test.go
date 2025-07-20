package main

import "fmt"

// 一个大而全的接口，强迫所有实现者都实现不需要的方法
type Worker interface {
	Work()
	Eat()
}

type HumanWorker struct{}

func (h HumanWorker) Work() {
	fmt.Println("Human is working")
}

func (h HumanWorker) Eat() {
	fmt.Println("Human is eating")
}

// 机器人不需要吃饭，但被强迫实现 Eat 方法
type RobotWorker struct{}

func (r RobotWorker) Work() {
	fmt.Println("Robot is working")
}

func (r RobotWorker) Eat() {
	// 不合理，机器人不吃饭
	fmt.Println("Robot doesn't need to eat, but forced to implement Eat()")
}

// 违反接口隔离原则的例子
func main() {
	var w Worker = HumanWorker{}
	w.Work()
	w.Eat()

	w = RobotWorker{}
	w.Work()
	w.Eat()
}

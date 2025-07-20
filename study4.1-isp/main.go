package main

import "fmt"

// 拆分接口：将职责细分
type Workable interface {
	Work()
}

type Eatable interface {
	Eat()
}

type HumanWorker struct{}

func (h HumanWorker) Work() {
	fmt.Println("Human is working")
}

func (h HumanWorker) Eat() {
	fmt.Println("Human is eating")
}

type RobotWorker struct{}

func (r RobotWorker) Work() {
	fmt.Println("Robot is working")
}

// RobotWorker 不实现 Eatable 接口，因为它不需要

func main() {
	var w Workable = HumanWorker{}
	w.Work()

	var e Eatable = HumanWorker{}
	e.Eat()

	w = RobotWorker{}
	w.Work()

	// RobotWorker 不实现 Eatable，所以不会被强制实现
}

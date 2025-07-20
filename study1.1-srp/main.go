package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 1. 订单实体：只负责“数据”
type Order struct {
	ID       int
	Customer string
	Items    []string
}

// 2. 订单操作：负责订单保存
type OrderSaver struct{}

func (os *OrderSaver) Save(order Order, filename string) error {
	file, err := os.createFile(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.MarshalIndent(order, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	return err
}

func (osa *OrderSaver) createFile(filename string) (*os.File, error) {
	return os.Create(filename)
}

// 3. 订单展示：负责订单展示
type OrderPrinter struct{}

func (op *OrderPrinter) Print(order Order) {
	fmt.Println("----- Order Details -----")
	fmt.Printf("Order ID: %d\n", order.ID)
	fmt.Printf("Customer: %s\n", order.Customer)
	fmt.Println("Items:")
	for _, item := range order.Items {
		fmt.Printf(" - %s\n", item)
	}
	fmt.Println("-------------------------")
}

// ----------- 主函数演示 ----------
// 目标：
//
// Order 结构体只负责“订单信息”的组织。
//
// 保存订单和打印订单由不同的结构体/模块负责，职责分离。

// 好处
// 易于维护和扩展（例如：换成数据库保存时只改 OrderSaver）
//
// 降低耦合，提高代码可读性和可测试性
func main() {
	order := Order{
		ID:       1001,
		Customer: "Alice",
		Items:    []string{"Apple", "Banana", "Orange"},
	}

	printer := OrderPrinter{}
	printer.Print(order)

	saver := OrderSaver{}
	err := saver.Save(order, "order.json")
	if err != nil {
		fmt.Println("保存失败:", err)
	} else {
		fmt.Println("订单已保存为 order.json")
	}
}

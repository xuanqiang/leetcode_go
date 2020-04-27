/**
 * 依赖倒置原则DIP(the Dependency Inversion Principle DIP)
 * 高层次的模块不应该依赖于低层次的模块，他们都应该依赖于抽象。
 * 抽象不应该依赖于具体实现，具体实现应该依赖于抽象。
 */
package main

import "fmt"

//抽象层
type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

//实现层
type BenZ struct {
}

func (benz *BenZ) Run() {
	fmt.Println("Benz is running...")
}

type Bmw struct {
	//...
}

func (bmw *Bmw) Run() {
	fmt.Println("Bmw is running...")
}

type Zhang_3 struct {
	//...
}

func (zhang3 *Zhang_3) Drive(car Car) {
	fmt.Println("Zhang3 drive car")
	car.Run()
}

type Li_4 struct {
	//...
}

func (li4 *Li_4) Drive(car Car) {
	fmt.Println("li4 drive car")
	car.Run()
}

//业务逻辑层
func main() {
	//张3 开 宝马
	var bmw Car
	bmw = &Bmw{}

	var zhang3 Driver
	zhang3 = &Zhang_3{}

	zhang3.Drive(bmw)

	//李4 开 奔驰
	var benz Car
	benz = &BenZ{}

	var li4 Driver
	li4 = &Li_4{}

	li4.Drive(benz)
}

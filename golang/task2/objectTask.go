package main

import "fmt"

func main() {
	rect := Rectangle{width: 10, heith: 12}
	circ := Circle{radius: 9}
	fmt.Println("矩形面积：%.2f,周长：%.2f", rect.Area(), rect.Perimeter())
	fmt.Println("圆的面积：%.2f,周长：%.2f", circ.Area(), circ.Perimeter())

	employee := Employee{name: "zhudongzhao", age: 18}
	person := Person{Employee: employee, employeeID: 12323}
	fmt.Println("用户信息：", person.PrintInfo())
}

// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width, heith float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.heith
}

func (r Rectangle) Perimeter() float64 {
	return (r.width + r.heith) * 2
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

// 使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
type Employee struct {
	name string
	age  int
}

type Person struct {
	Employee
	employeeID int
}

func (p Person) PrintInfo() string {
	return fmt.Sprintf("Name: %s, Age: %d, EmployeeID: %d", p.name, p.age, p.employeeID)
}

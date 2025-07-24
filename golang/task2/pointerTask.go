package main

import (
	"fmt"
)

func main() {
	var p int = 16
	sumPoint(&p)
	fmt.Println("p = %d", p)

	var pp []int = []int{21, 12, 3, 10}
	for v := range pp {
		mulPoint(&pp[v])
	}
	fmt.Println("pp = %d", pp)
}

// 定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
func sumPoint(num *int) {
	*num += 10
}

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
func mulPoint(num *int) {
	*num *= 2
}

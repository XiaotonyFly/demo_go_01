package main

import (
	"fmt"
	"math"
)

func main() {
	//函数作为实参
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(16))
	testFunc()

	fmt.Println("-----------------分割线-----------------")
	aa := getBiBao()
	fmt.Println(aa())
	fmt.Println(aa())
	fmt.Println(aa())
	fmt.Println("-----------------分割线-----------------")
	tea1 := Tea{"红茶", "福建"}
	tea1.aa()
}

type cb func(int) int

func testCallBack(x int, f cb) {
	f(x)
}

func callBack(x int) int {
	fmt.Printf("我是回调call，x：%d\n", x)
	return x
}

func testFunc() {
	testCallBack(2, callBack)
	testCallBack(5, func(i int) int {
		fmt.Printf("woshihuidiao: x: %d\n", i)
		return i
	})
}

func getBiBao() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

type Tea struct {
	name        string
	growingArea string
}

func (tea Tea) aa() {
	fmt.Printf("茶的名称为%s，产地为%s\n", tea.name, tea.growingArea)
}

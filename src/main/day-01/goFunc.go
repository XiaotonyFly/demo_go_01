package main

import (
	"fmt"
	"math"
)

func main() {
	/*
	   Go 语言函数定义格式如下：
	   	func function_name( [parameter list] ) [return_types] {
	   	   函数体
	   	}

	   函数定义解析：

	   func：函数由 func 开始声明
	   function_name：函数名称，函数名和参数列表一起构成了函数签名。
	   parameter list：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
	   return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
	   函数体：函数定义的代码集合。
	*/

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

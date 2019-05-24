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

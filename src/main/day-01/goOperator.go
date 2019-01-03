package main

import "fmt"

func main() {
	showArithmeticOperator()
	showGxOperator()
}

//算术运算符
//相加（+）、相减（-）、相乘（*）、相除（/）、求余（%）、自增（++）、自减（--）
const x1, y1 = 12, 180

func showArithmeticOperator() {
	fmt.Println(x1 + y1)
	fmt.Println(y1 - x1)
	fmt.Println(y1 * x1)
	fmt.Println(y1 / x1)
	fmt.Println(y1 % x1)
	var a = 3
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)
}

//关系运算符
//等于（==）、不等于（!=）、大于（>）、小于（<）、大于等于（>=）、小于等于（<=）
func showGxOperator() {
	a, b := 3, 4
	if a == b {
		fmt.Println("a 等于 b")
	} else {
		if a > b {
			fmt.Println("a 大于 b")
		} else {
			fmt.Println("a 小于 b")
		}
		fmt.Println("a 不等于 b")
	}

	var p *int
	p = &a
	fmt.Println(p)
	fmt.Println(&a)
	fmt.Println(*p)
}

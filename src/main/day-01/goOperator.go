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
//Go语言没有三目运算符（即不支持?:形式的条件判断）
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

	/*
		Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。
		指针声明：
			var variable_name *variable_type
		variable_type 为指针类型，variable_name 为指针变量名，* 号用于指定变量是作为一个指针。
	*/
	var p *int
	p = &a
	fmt.Println(p)
	fmt.Println(&a)
	//在指针变量前面加上 * 号（前缀）来获取指针所指向的内容。
	fmt.Println(*p)
}

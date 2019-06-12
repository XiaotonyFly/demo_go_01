package main

import (
	"fmt"
)

func main() {
	/*
	   数组的声明:
	   	var variable_name  [SIZE] variable_type
	   多维数组声明：var variable_name  [SIZE1][SIZE2]...[SIZEN] variable_type

	   数组初始化：
	   var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	   初始化数组中 {} 中的元素个数不能大于 [] 中的数字。

	   如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：

	    var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	   或
	    var balance = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	*/
	var intArray [3]int
	fmt.Println(len(intArray))
	fmt.Println(intArray[2])

	s := []string{"aa", "bb"}
	fmt.Println(s[1])

	a := "bnnb"
	var ptr *string
	ptr = &a
	fmt.Printf("%x\n", ptr)
	fmt.Println(ptr)
	fmt.Printf("%s\n", *ptr)
	fmt.Printf("%x\n", &ptr)
	fmt.Printf("%x\n", &a)

	var p *int
	b := 15
	p = &b
	if p == nil {
		fmt.Println("p 为空指针")
	} else {
		fmt.Println("p 不为空指针")
	}
	ads(&b)
	fmt.Println("---------------分割线-------------")

	var balance = [...]float32{1, 2, 6.9, 3.4}
	fmt.Println("balance数组长度为：", len(balance))
}

func ads(h *int) {
	fmt.Println(h)
	fmt.Println(*h)
}

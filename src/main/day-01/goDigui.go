package main

import "fmt"

func main() {
	//DiGui(5)
	cheng := JieCheng(4)
	fmt.Println(cheng)

	for i := 0; i < 10; i++ {
		fmt.Printf("%d\t", Fibonacci(i))
	}

	fmt.Println()
	var a int
	a = 4
	var b float32
	fmt.Printf("a = %f\n", float32(a))
	fmt.Printf("b = %f\n", b)
}

func DiGui(i int) {
	if i >= 0 {
		fmt.Printf("递归函数调用：%d\n", i)
		i--
		DiGui(i)
	} else {
		fmt.Println("退出递归函数")
	}
}

func JieCheng(num uint64) uint64 {
	if num > 0 {
		result := num * JieCheng(num-1)
		return result
	}
	return 1
}

func Fibonacci(n int) int {
	if n > 2 {
		return Fibonacci(n-1) + Fibonacci(n-2)
	} else {
		return n
	}
}

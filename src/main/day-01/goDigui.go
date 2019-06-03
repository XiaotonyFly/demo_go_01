package main

import "fmt"

func main() {
	//DiGui(5)
	cheng := JieCheng(4)
	fmt.Println(cheng)
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

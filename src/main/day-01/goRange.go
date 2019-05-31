package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 5}
	for _, a := range arr {
		fmt.Println(a)
	}

	nums := []int{1, 2, 3, 4}
	length := 0
	for range nums {
		length++
	}
	fmt.Println(length)
	for i, v := range nums {
		fmt.Printf("第%d的元素为%d\n", i, v)
	}
}

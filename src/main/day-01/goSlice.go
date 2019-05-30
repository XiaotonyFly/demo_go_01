package main

import "fmt"

func main() {
	var numbers []int
	printSlice(numbers)

	numbers = append(numbers, 0)
	printSlice(numbers)

	numbers = append(numbers, 1)
	printSlice(numbers)

	numbers = append(numbers, 2)
	printSlice(numbers)

	numbers = append(numbers, 3, 4, 5, 6)
	printSlice(numbers)

	//twoDimensionArray()
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func twoDimensionArray() {
	/* 数组 - 5 行 2 列*/
	var a = [][]int{{0, 0}, {1, 2}, {2}, {3, 6}, {4, 8}}
	var i, j int

	/* 输出数组元素 */
	for i = 0; i < len(a); i++ {
		for j = 0; j < len(a[i]); j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}

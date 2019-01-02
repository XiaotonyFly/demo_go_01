package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("你好")
	fmt.Println(math.Pi)

	var a int
	a = 2
	fmt.Printf("a is %T\n", a)
	fmt.Println(a)
	p := &a
	fmt.Println(&a)
	fmt.Println(*p)
}

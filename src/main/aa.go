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

	var b = 3.9
	fmt.Println(b)

	p := &a
	fmt.Println(&a)
	fmt.Println(*p)

	fmt.Println(Ss(float32(a), float32(b)))

}

func Ss(a, b float32) float32 {
	var c = a * b
	return c
}

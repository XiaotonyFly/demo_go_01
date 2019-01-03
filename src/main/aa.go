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
	switchTest()
}

func Ss(a, b float32) float32 {
	var c = a * b
	return c
}

func switchTest() {
	grade := "D"
	marks := 96

	switch {
	case (100 >= marks && marks >= 90):
		grade = "A"
	case (90 > marks && marks >= 80):
		grade = "B"
	case (80 > marks && marks >= 60):
		grade = "C"
	default:
		grade = "D"
	}

	switch grade {
	case "A":
		fmt.Println("成绩优秀")
	case "B":
		fmt.Println("成绩良好")
	case "C":
		fmt.Println("成绩及格")
	default:
		fmt.Println("成绩不及格")
	}

	fmt.Printf("你的成绩是 %s\n", grade)

	var m interface{}
	switch i := m.(type) {
	case nil:
		fmt.Printf("m的类型 :%T", i)
	case string:
		fmt.Println("m的类型是string")
	}
}

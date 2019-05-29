package main

import (
	"fmt"
)

func main() {
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
}

func ads(h *int) {
	fmt.Println(h)
	fmt.Println(*h)
}

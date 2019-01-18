package main

import "fmt"

func main() {
	//Go语言的For循环有3中形式，只有其中的一种使用分号。
	/*
		第一种：for init; condition; post { }
		init： 一般为赋值表达式，给控制变量赋初值；
		condition： 关系表达式或逻辑表达式，循环控制条件；
		post： 一般为赋值表达式，给控制变量增量或减量。
	*/
	for i := 0; i < 100; i++ {
		if i%3 == 0 && i > 0 {
			fmt.Print(i)
		}
	}
	fmt.Println("")
	getChengFaBiao()

	/*
		第二种：for condition {}
	*/
	b := 15
	a := 1
	for a < b {
		fmt.Printf(" a = %d ", a)
		a++
	}

	/*
		第三种：for {},相当于C语言的for(;;)
	*/
	//无限循环（死循环）
	/*for {
		fmt.Println("aaa")
	}*/

	//for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。
	array := [6]int{1, 3, 6, 7}
	for i, value := range array {
		fmt.Printf("第 %d 个元素为 %d\n", i, value)
	}

	getSuShu()

}

/*
	Go 语言嵌套循环的格式：

	for [condition |  ( init; condition; increment ) | Range]{
   		for [condition |  ( init; condition; increment ) | Range]{
      		statement(s);
   		}
   		statement(s);
	}
*/
//乘法表
func getChengFaBiao() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf(" %d * %d = %d ", j, i, i*j)
		}
		fmt.Println("")
	}
}

//素数
func getSuShu() {
	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break // 如果发现因子，则不是素数
			}
		}
		if j > (i / j) {
			fmt.Println(j)
			fmt.Printf("%d  是素数\n", i)
		}
	}
}

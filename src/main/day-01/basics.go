package main

import "fmt"

func main() {
	fmt.Println(globalVariable)
	fmt.Println(x, y)
	getJbVariable()
	getMultivariable()
	fmt.Println(CL)
	fmt.Println(CL_1)
	getConst()
}

//Go:变量
//1、变量的声明
var globalVariable int //(全局变量允许声明但不使用，局部变量声明了就一定要使用，不然会报错)
var (                  //这种因式分解关键字的写法一般用于声明全局变量
	x int
	y float64
)

func getJbVariable() {
	//局部变量
	/*第一种方式：指定变量类型，声明后若不赋值，使用默认值。
	var v_name v_type
	v_name = value
	*/
	var a int
	a = 2
	fmt.Println("a === ", a)
	/*第二种方式：根据值自行判定变量类型
	var v_name = value
	*/
	var b = 3.6
	fmt.Println("b === ", b)
	fmt.Printf("b is %T\n", b)  //获取b的类型
	fmt.Printf("b is %p\n", &b) //获取b的内存地址

	/*第三种方式：省略var, 注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误（不适用于全局变量的声明, 即只能在函数体中运用）。
	v_name := value
	*/
	c := "ssdfd"
	fmt.Println(c)
	fmt.Printf("c is %T \n", c)

}

func getMultivariable() {
	//多变量声明
	var d, e, f string
	d, e, f = "ddd", "eee", "fff"
	fmt.Println(d, e, f)

	var a, b, c = 1, 2.4, "cc"
	fmt.Println(a, b, c)

	g, h := 1, 3
	fmt.Println(g, h)
}

/*空白标识符(下划线) _ 也被用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃。

_ (下划线)实际上是一个只写变量，你不能得到它的值。这样做是因为 Go 语言中你必须使用所有被声明的变量，但有时你并不需要使用从一个函数得到的所有返回值。

并行赋值也被用于当一个函数返回多个返回值时，比如这里的 val 和错误 err 是通过调用 Func1 函数同时得到：val, err = Func1(var1)。*/

//2、Go语言常量 (格式：const identifier [type] = value)
const CL string = "这是一个常量" //显式定义
const CL_1 = "这是一个常量1"     //隐式定义

//每左移一位相当于乘以2的1次方，每右移一位相当于除以2的1次方
const (
	u = 1 << iota
	l = 3 << iota
	j
	k
)

//iota，特殊常量，可以认为是一个可以被编译器修改的常量。

//iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
func getConst() {
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
	fmt.Println(u, l, j, k)
}

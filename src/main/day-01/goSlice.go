package main

import (
	"fmt"
	"unsafe"
)

/*
	1、Go 语言切片是对数组的抽象。
		Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),
	与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

	2、定义切片
		可以声明一个未指定大小的数组来定义切片（切片不需要说明长度）：
			var identifier []type
		或使用make()函数来创建切片:
			var slice1 []type = make([]type, len)
		也可以简写为 slice1 := make([]type, len)
		也可以指定容量，其中capacity为可选参数。
		make([]T, length, capacity)
		这里 len 是数组的长度并且也是切片的初始长度。

	3、切片初始化
		s :=[] int {1,2,3 } ——》 直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3。
		s := arr[:] ——》 初始化切片s,是数组arr的引用。
		s := arr[startIndex:endIndex] ——》 将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片。
		s := arr[startIndex:] ——》 默认 endIndex 时将表示一直到arr的最后一个元素。
		s := arr[:endIndex] ——》 默认 startIndex 时将表示从arr的第一个元素开始。
		s1 := s[startIndex:endIndex] ——》 通过切片s初始化切片s1。
		s :=make([]int,len,cap) ——》 通过内置函数make()初始化切片s,[]int 标识为其元素类型为int的切片。

		len() 和 cap() 函数
		切片是可索引的，并且可以由 len() 方法获取长度。
		切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。

		空(nil)切片，一个切片在未初始化之前默认为 nil，长度为 0。

		切片截取，可以通过设置下限及上限来设置截取切片 [lower-bound:upper-bound]（不包含upper-bound索引）。

		注：(我们基于原数组或者切片创建一个新的切片后，那么新的切片的大小和容量是多少呢？
			这里有个公式，对于底层数组容量是 k 的切片 slice[i:j] 来说：长度: j-i,容量: k-i)。

		append() 和 copy() 函数
         如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。
*/
func main() {
	/*var numbers []int
	printSlice(numbers)

	numbers = append(numbers, 0)
	printSlice(numbers)

	numbers = append(numbers, 1)
	printSlice(numbers)

	numbers = append(numbers, 2)
	printSlice(numbers)

	numbers = append(numbers, 3, 4, 5, 6)
	printSlice(numbers)*/

	//twoDimensionArray()

	slice_test := []int{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(slice_test))
	fmt.Printf("main:%#v,%#v,%#v\n", slice_test, len(slice_test), cap(slice_test))
	slice_value(slice_test)
	fmt.Printf("main:%#v,%#v,%#v\n", slice_test, len(slice_test), cap(slice_test))
	slice_ptr(&slice_test)
	fmt.Printf("main:%#v,%#v,%#v\n", slice_test, len(slice_test), cap(slice_test))
	fmt.Println(unsafe.Sizeof(slice_test))
}

func printSlice(x []int) {
	if x == nil {
		fmt.Println("切片为空")
	}
	fmt.Printf("len=%d cap=%d slice=%v type=%T\n", len(x), cap(x), x, x)
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

func slice_value(slice_test []int) {
	slice_test[1] = 100                // 函数外的slice确实有被修改
	slice_test = append(slice_test, 6) // 函数外的不变
	fmt.Printf("slice_value:%#v,%#v,%#v\n", slice_test, len(slice_test), cap(slice_test))
}

func slice_ptr(slice_test *[]int) { // append要传指针才能修改函数外的slice
	*slice_test = append(*slice_test, 7)
	fmt.Printf("slice_ptr:%#v,%#v,%#v\n", *slice_test, len(*slice_test), cap(*slice_test))
}

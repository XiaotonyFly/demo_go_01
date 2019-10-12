package main

import "fmt"

func main() {
	/*goLang：
		Map 是一种无序的键值对的集合。
		Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
		Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，
			我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。

	1、定义 Map
	可以使用内建函数 make 也可以使用 map 关键字来定义 Map:
	声明变量，默认 map 是 nil
		var map_variable map[key_data_type]value_data_type

	使用 make 函数
		map_variable := make(map[key_data_type]value_data_type)
	如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对

	2、delete() 函数
	delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key。
	*/
	var goMap map[string]string     //定义map
	goMap = make(map[string]string) //初始化map
	goMap["apple"] = "苹果"
	goMap["orange"] = "橙子"
	goMap["banana"] = "香蕉"
	for fruit := range goMap {
		fmt.Println(fruit + "=" + goMap[fruit])
	}

	s, ok := goMap["a"]
	if ok {
		fmt.Println(s)
	} else {
		fmt.Println(ok)
	}

	for fruit := range goMap {
		fmt.Println(fruit, "中文名称是：", goMap[fruit])
	}

	fmt.Println("删除Apple")
	delete(goMap, "apple")

	for fruit := range goMap {
		fmt.Println(fruit, "中文名称是：", goMap[fruit])
	}
}

package main

import "fmt"

func main() {
	/*goLang：
	Map 是一种无序的键值对的集合。
	Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
	Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。
	*/
	var goMap map[string]string
	goMap = make(map[string]string)
	goMap["apple"] = "苹果"
	goMap["orange"] = "橙子"
	for fruit := range goMap {
		fmt.Println(fruit + "=" + goMap[fruit])
	}

	s, ok := goMap["a"]
	if ok {
		fmt.Println(s)
	} else {
		fmt.Println(ok)
	}

}
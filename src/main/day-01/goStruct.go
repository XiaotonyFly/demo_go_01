package main

import "fmt"

/*
定义结构体
	结构体定义需要使用 type 和 struct 语句。struct 语句定义一个新的数据类型，结构体有中有一个或多个成员。type 语句设定了结构体的名称。结构体的格式如下：

	type struct_variable_type struct {
		member definition;
		member definition;
		...
		member definition;
	}

一旦定义了结构体类型，它就能用于变量的声明，语法格式如下：

variable_name := structure_variable_type {value1, value2...valuen}
或
variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}

访问结构体成员
如果要访问结构体成员，需要使用点号 . 操作符，格式为：

结构体.成员名"
*/
type Animal struct {
	name    string
	country string
}

func main() {
	var pander Animal
	var dog Animal
	pander.name = "熊猫"
	pander.country = "中国"
	dog.name = "狗"
	dog.country = "地球"
	cat := Animal{"猫", "中国"}
	fmt.Println(dog.country)
	printAnimal(dog)
	printAnimalByPointer(&dog)
	printAnimal(cat)
}

func printAnimal(animal Animal) {
	fmt.Println(animal.name)
}

func printAnimalByPointer(animal *Animal) {
	/*
		结构体指针
		你可以定义指向结构体的指针类似于其他指针变量，格式如下：
			var struct_pointer *Books

		以上定义的指针变量可以存储结构体变量的地址。查看结构体变量地址，可以将 & 符号放置于结构体变量前：
			struct_pointer = &Book1;
		使用结构体指针访问结构体成员，使用 "." 操作符：
			struct_pointer.title;
	*/
	fmt.Println(animal.name)
}

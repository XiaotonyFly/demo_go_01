package main

import "fmt"

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
	fmt.Println(dog.country)
	printAnimal(dog)
	printAnimalByPointer(&dog)
}

func printAnimal(animal Animal) {
	fmt.Println(animal.name)
}

func printAnimalByPointer(animal *Animal) {
	fmt.Println(animal.name)
}

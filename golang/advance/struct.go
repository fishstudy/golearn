package main

import "fmt"

type Dog struct {
	ID   int
	Name string
	Age  int
}

func testForStruct() {
	var dog Dog
	dog.ID = 1
	dog.Age = 3
	dog.Name = "旺财"
	fmt.Println(dog)
}

func testForStruct() {
	dog := Dog{}
}

func main() {
	testForStruct()
}

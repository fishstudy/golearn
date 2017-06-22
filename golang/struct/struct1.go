package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {
	a := &person{
		Name: "yuxuefeng",
		Age:  30,
	}
	a.Name = "sunyanyan"
	fmt.Println(a)
	A(a)
	B(a)
	fmt.Println(a)
}

func A(per *person) {
	per.Age = 19
	fmt.Println(per)
}

func B(per *person) {
	per.Age = 20
	fmt.Println("A", per)
}

package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

type person1 struct {
	Name string
	Age  int
}

func main() {
	a := person{Name: "yuxuefeng", Age: 33}
	b := person{Name: "yuxuefeng", Age: 33}
	fmt.Println(b == a)
}

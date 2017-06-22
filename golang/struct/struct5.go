package main

import (
	"fmt"
)

type human struct {
	Sex int
}

type teacher struct {
	human
	Name string
	Age  int
}

type student struct {
	human
	Name string
	Age  int
}

func main() {
	a := teacher{Name: "yuxuefeng", Age: 33, human: human{Sex: 0}}
	b := student{Name: "yuxuefeng", Age: 33, human: human{Sex: 1}}
	a.Name = "yufeng"
	a.Age = 16
	a.Sex = 10
	fmt.Println(a, b)
}

package main

import (
	"fmt"
)

type A struct {
	Name string
}

type B struct {
	Name string
}

func main() {
	a := A{}
	a.Print()
	fmt.Println(a.Name)

	b := B{}
	b.Print()
	fmt.Println(b.Name)
}

func (a *A) Print() {
	a.Name = "AA"
	fmt.Println("a")
}

func (b B) Print() {
	b.Name = "BB"
	fmt.Println("b")
}

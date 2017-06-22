package main

import (
	"fmt"
)

type A struct {
	B
	Name string
}

type B struct {
	Name string
}

func main() {
	a := A{Name: "A", B: B{Name: "B"}}
	fmt.Println(a.Name, a.B.Name)
}

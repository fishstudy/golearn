package main

import (
	"fmt"
)

type A struct {
	name string
}

func main() {
	a := A{}
	a.Print()

}

func (a *A) Print() {
	a.name = "123"
	fmt.Println(a.name)
}

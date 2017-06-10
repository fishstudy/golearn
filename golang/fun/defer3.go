package main

import (
	"fmt"
)

func main() {
	A()
	B()
	C()
}

func A() {
	fmt.Println("Func A")
}

func B() {
	panic("Panic in B")
}

func C() {
	fmt.Println("Func C")
}

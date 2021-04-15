package main

import (
	"fmt"
	"math"
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Println(a)
	fmt.Printf("%d %q", a, s)
}
func variableInit() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func shortInit() {
	a, b := 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func traingle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInit()
	traingle()
}

package main

import (
	"fmt"
)

func main() {
	a := [10]int{}
	a[1] = 2
	fmt.Println(a)
	p := new([10]int)
	p[1] = 2
	fmt.Println(p)
}

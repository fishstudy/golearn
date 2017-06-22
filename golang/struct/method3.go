package main

import (
	"fmt"
)

type TZ int

func (a *TZ) Increase(num int) {
	*a += TZ(num)
}

func main() {
	var a TZ
	a.Increase(100)
	fmt.Println(a)
}

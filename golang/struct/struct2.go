package main

import (
	"fmt"
)

func main() {
	a := &struct {
		Name string
		Age  int
	}{
		Name: "yuxuefeng",
		Age:  30,
	}
	fmt.Println(a)
}

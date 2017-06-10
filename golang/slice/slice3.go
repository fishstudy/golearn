package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	s1 := a[2:5]
	s2 := a[1:3]
	fmt.Println(s1, s2)
	s1[0] = 9
	fmt.Println(s1, s2)
}

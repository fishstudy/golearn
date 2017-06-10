package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3, 4}
	A(s1)
	fmt.Println(s1)
}

func A(s []int) {
	s[0] = 5
	s[1] = 6
	s[2] = 7
	s[3] = 8
	fmt.Println(s)
}

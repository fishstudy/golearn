package main

import "fmt"

func main() {
	ary := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := ary[2:9]
	fmt.Println(cap(s))
	fmt.Println(s)
	s3 := append(s, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	s2 := s[3:8]
	fmt.Println(s2)
	fmt.Println(cap(s2))
	fmt.Println(s3, s4, s5)
	fmt.Println(ary)
}

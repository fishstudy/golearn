package main

import "fmt"

func main() {

	s1 := []int{2, 4, 6, 8, 10}
	fmt.Println(s1)
	s2 := make([]int, 16)
	s3 := make([]int, 32)
	fmt.Println(s2, s3)
	copy(s2, s1)

	len := len(s2)
	s2[len-1] = 100
	fmt.Println(s2)
	//head := s2[0]
	s2 = s2[1:]
	fmt.Println(s2)

	//len2 := len(s2)
	//s2 = s2[:len2-1]
	//fmt.Println(s2)
}

package main

import "fmt"

func main() {
	var ary1 [5]int
	ary2 := [3]int{1, 3, 5}
	ary3 := [...]int{2, 4, 6, 8, 10}
	fmt.Println(ary1, ary2, ary3)
}

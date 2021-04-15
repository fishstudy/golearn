package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str = "hello 世界"
	fmt.Println("len(str):", len(str))
	fmt.Println("RuneCountInString:", utf8.RuneCountInString(str))
	fmt.Println("rune:", len([]rune(str)))
}

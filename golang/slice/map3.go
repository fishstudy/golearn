package main

import (
	"fmt"
)

func main() {
	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
	s := make(map[string]int)

	for k, v := range m {
		s[v] = k
	}
	fmt.Println(m)
	fmt.Println(s)
}

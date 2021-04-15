package main

import "fmt"

func main() {
	m := map[string]string{
		"name":   "ccmouse",
		"course": "goland",
		"site":   "kona",
	}
	fmt.Println(m)

	m2 := make(map[string]int)
	fmt.Println(m2)
	for _, v := range m {
		fmt.Println(v)
	}
}

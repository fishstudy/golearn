package main

import (
	"fmt"
)

func main() {
	a := [...]int{5, 2, 7, 9, 1}
	t := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k < 5; k++ {
				if a[i] != a[j] && a[j] != a[k] && a[i] != a[k] {
					t++
					fmt.Println(a[i], a[j], a[k])
				}
			}
		}
	}
	fmt.Println(t)
}

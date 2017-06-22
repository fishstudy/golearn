package main

import (
	"fmt"
)

func main() {
	var I float32 = 0.0
	var bonus float32 = 0.0
	fmt.Println("Enter the lirun")
	fmt.Scanf("%fn", &I)
	switch {
	case I > 1000000:
		bonus = (I - 1000000) * 0.01
		fallthrough
	case I > 6000000:
		bonus += (I - 6000000) * 0.015
		I = 6000000
		fallthrough
	case I > 4000000:
		bonus += (I - 4000000) * 0.03
		I = 4000000
		fallthrough
	case I > 2000000:
		bonus += (I - 2000000) * 0.05
		fallthrough
	case I > 100000:
		bonus += (I - 100000) * 0.75
		I = 100000
		fallthrough
	default:
		bonus += I * 0.1
	}
	fmt.Printf("get the bill%fn", bonus)
}

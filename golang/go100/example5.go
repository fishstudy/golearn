package main

import (
	"fmt"
)

func tao(day int) (sum int) {
	if day == 10 {
		return 1
	} else {
		sum = (tao(day+1) + 1) * 2
		fmt.Println(sum)
		return sum
	}
}

func main() {
	var total int = 0
	total = tao(1)
	fmt.Printf("count=%d\n", total)
}

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 21:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

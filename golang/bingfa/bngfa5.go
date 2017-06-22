package main

import (
	"fmt"
)

func main() {
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool, 2)
	go func() {
		a, b := false, false
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					fmt.Println("C1C1C1")
					if !a {
						o <- true
						a = true
					}
					break
				}
				fmt.Println("C1", v)
			case v, ok := <-c2:
				if !ok {
					fmt.Println("C2C2C2")
					if !b {
						o <- true
						b = true
					}

					break
				}
				fmt.Println("C2", v)
			}
		}
	}()
	c1 <- 1
	c2 <- "HI"
	c1 <- 3
	c2 <- "Hello"

	close(c1)
	close(c2)
	for i := 0; i < 2; i++ {
		<-o
	}
}

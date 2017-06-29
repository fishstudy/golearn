package main

import "fmt"

func main() {
	message := make(chan string)

	go func() {
		fmt.Println("recive ping")
		message <- "ping"
	}()
	fmt.Println("after go")
	msg := <-message
	fmt.Println(msg)
}

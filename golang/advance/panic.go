package main

import "fmt"

func main() {
	receivePanic()
	//defer coverPanic()
}

func receivePanic() {
	defer coverPanic()
	panic("I am panic")
}

func coverPanic() {
	message := recover()
	fmt.Println(message)
}

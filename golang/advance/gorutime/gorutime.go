package gorutime

import (
	"fmt"
	"time"
)

func Loop() {
	for i := 1; i < 11; i++ {
		fmt.Printf("%d,", i)
		time.Sleep(time.Microsecond * 10)
	}
}

var ch chan int = make(chan int, 5)
var timeout chan bool = make(chan bool, 10)

func Send() {
	time.Sleep(time.Second * 1)
	ch <- 1
	timeout <- false
	time.Sleep(time.Second * 1)
	ch <- 2
	timeout <- true
	time.Sleep(time.Second * 1)
	ch <- 3
	time.Sleep(time.Second * 1)
	ch <- 4
	time.Sleep(time.Second * 1)
	timeout <- true

}

func Reciver() {
	for {
		select {
		case num := <-ch:
			fmt.Printf("num=%d", num)
			fmt.Println("")
		case <-timeout:
			fmt.Println("timeout ....")
		}
	}
}

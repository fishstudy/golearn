package main

import "fmt"

func main() {
	str:=" I am 于雪锋"
	fmt.Println(len(str))
	s := []rune(str)
	 for _,val := range s {
		fmt.Printf("%v",val)
	}
}

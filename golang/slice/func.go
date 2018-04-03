package main

import "fmt"

func main(){
	var (
		a = 10
		b = 100
	)

	a2, b2 :=  add(a, b)
	fmt.Println("a2, b2=",a2, ", ", b2)

}

func add(n1, n2 int)(res1, res2 int) {
	res1, res2 = n1*n2, n2*2
	return
}

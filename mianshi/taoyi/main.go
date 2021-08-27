package main

import "fmt"

type a struct {

}

func NewA() *a {
	A:=new(a)
	fmt.Println("指针逃逸")
	return A
}

func closure()func()  {
	closer:="闭包逃逸"
	return func() {
		closer="逃逸成功"
		fmt.Println(closer)
	}
	
}
func main() {
	NewA()
	slice := make([]int,100,100)
	fmt.Println("栈空间不足逃逸",slice)
	strings:="动态逃逸"
	fmt.Printf("%v",strings)
	closure()
}

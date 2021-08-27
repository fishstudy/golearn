package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	var s1 []int   // nil切片
	s:=new([]string)
	fmt.Println(s1)
	fmt.Println(s)
	//*s[0] = "egg"
	*s = append(*s,"egg")
	*s = append(*s,"dog")
	*s[2] = "cat"
	fmt.Println(s)
	s2 := make([]int,0)  // 空切片
	s4 := make([]int,0)   // 空切片

	fmt.Printf("s1 pointer:%+v, s2 pointer:%+v, s4 pointer:%+v, \n", *(*reflect.SliceHeader)(unsafe.Pointer(&s1)),*(*reflect.SliceHeader)(unsafe.Pointer(&s2)),*(*reflect.SliceHeader)(unsafe.Pointer(&s4)))
	fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s1))).Data==(*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data)
	fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data==(*(*reflect.SliceHeader)(unsafe.Pointer(&s4))).Data)
}

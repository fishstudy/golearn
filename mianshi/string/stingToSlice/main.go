package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a :="aaa"
	ssh := *(*reflect.StringHeader)(unsafe.Pointer(&a))
	b := *(*[]byte)(unsafe.Pointer(&ssh))
	fmt.Printf("%v",b)
}

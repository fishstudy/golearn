package main

import (
	"fmt"
	"reflect"
)

type T1 struct {
	Name string
	Age int
	Arr [2]bool
	ptr *int
	//slice []int
	//map1 map[string]string
}

type T2 struct {
	Name string
	Age int
	Arr [2]int
	ptr *int
	slice []int
	map1 map[string]string
}

func main() {
	t1 := T1{
		Name:"111",
		Age:8,
		Arr:[2]bool{true,true},
		ptr:new(int),
		//slice:[]int{1,2,3},
		//map1:make(map[string]string),
	}
	t2:=T1{
		Name:"111",
		Age:8,
		Arr:[2]bool{true,true},
		ptr:new(int),
		//slice:[]int{1,2,3},
		//map1:make(map[string]string),
	}
	fmt.Println(t1 == t2)
	fmt.Println(reflect.DeepEqual(t1 ,t2))
	fmt.Printf("%p,%p",t1,t2)
}

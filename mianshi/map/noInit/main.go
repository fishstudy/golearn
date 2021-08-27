package main

import "fmt"

func main() {
	var m map[string]string
	m = make(map[string]string, 10)
	fmt.Println(m)
	m["a"] = "wangwu"
	m["b"] = "zhangsan"
	fmt.Println(m)

	m2 := make(map[string]string, 3)
	m2["a"] = "a"
	m2["b"] = "b"
	m2["c"] = "d"
	m2["d"] = "d"
	fmt.Println(m2)


	var m3 map[string]string
	m3 = map[string]string{
		"n1" : "x",
		"n2" : "y",
		"n3" : "z",//最后一个逗号不能去掉
	}
	m3["a"] = "a"
	fmt.Println(m3)


}

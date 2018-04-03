package main

import "fmt"

func main(){
  uncert(1, 3, 5)
  uncert(2, 4)
  //3.2 不同类型
  uncert2(10.01, "Hello")
  //uncert3("hello", 1)
 // uncert3("hello", 1, 2, 3)
}

func uncert(args ...int) {
  for k, v := range args {
    fmt.Printf("%d => %d\n", k, v)
  }
}

func uncert2(args ...interface{}) {
  for k, v := range args {
    fmt.Println(k, " => ", v)
  }
  
}

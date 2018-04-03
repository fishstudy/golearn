package main

import "fmt"

func main(){
 s1 := make([]int, 3, 6)
 fmt.Printf("%p\n",s1);

s1 = append(s1,1,2,3)
fmt.Printf("%v %p %v\n",s1, s1,cap(s1))

s1 = append(s1,1,2,3)
fmt.Printf("%v %p %v\n",s1, s1,cap(s1))
}

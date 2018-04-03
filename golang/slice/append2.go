package main

import "fmt"

func main(){
var a = [5]int{1,2,3,4,5}

var s1, s2 []int
s1 =  a[1:4]
s2 =  a[2:4]
fmt.Printf("%v %p %v\n",s1, s1,cap(s1))
fmt.Printf("%v %p %v\n",s2, s2,cap(s2))
s2 = append(s2,11,12,13, 14,15,16,17)
s1[1] = 100
fmt.Printf("%v %p %v\n",s1, s1,cap(s1))
fmt.Printf("%v %p %v\n",s2, s2,cap(s2))

}

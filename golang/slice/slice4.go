package main

import "fmt"

func main(){
   var a = [10]int {1,2,3,4,5,6,7,8,9,10}
   var s, s2,s3 []int
   s = a[2:6]
   s2 =  a[3:5]
   s3 = s[1:5]
   s[1] = 33
   fmt.Println(s)
   fmt.Println(s2)
   fmt.Println(s3)
   fmt.Println(a)

}

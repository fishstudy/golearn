package main

import "fmt"

func main() {

	src := "你好啊,XXX啊啊啊啊"
	dst := reverse([]rune(src))
	fmt.Println("result is:",string(dst))
}

func reverse(s []rune) []rune {
	for i,j :=0,len(s)-1;i<j;i,j = i+1,j-1{
		s[i],s[j] = s[j],s[i]
	}
	return s
}



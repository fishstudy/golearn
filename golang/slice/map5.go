package main

import(
	"fmt"
) 

func main() {
    var m = map[string]int{
        "unix":         0,
        "python":       1,
        "go":           2,
        "javascript":   3,
        "testing":      4,
        "philosophy":   5,
        "startups":     6,
        "productivity": 7,
        "hn":           8,
        "reddit":       9,
        "C++":          10,
    }
    var  m2 = make(map[int]string)
    fmt.Println(m)
    for k,v := range m {
       m2[v] = k
    }
    fmt.Println(m2)
}

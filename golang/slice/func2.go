package main

import (
	"fmt"
	"math"
	"strconv"
)

func main(){
        b := 15
	fmt.Printf("sqrt(%d) = %.2f\n", b, math.Sqrt(float64(b)))
	c := 67
	fmt.Println("string(67)=", string(c))
	fmt.Println("strconv.Itoa(67)=", strconv.Itoa(c))
}


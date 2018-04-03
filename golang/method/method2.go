package main

import (
	"fmt"
	"math"
)

type Rectangel struct {

	width,height float64
}

type Circle struct {
	radius float64
}



func (r Rectangel) area() float64{
	return r.width * r.height
}


func (c Circle) area() float64{
	return c.radius * c.radius * math.Pi
}

func main() {
        r1 := Rectangel{12, 2}
        r2 := Rectangel{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}
	
	fmt.Println("Area of r1 is:", r1.area())
	fmt.Println("Area of r2 is:", r2.area())
	fmt.Println("Area of c1 is:", c1.area())
	fmt.Println("Area of c2 is:", c2.area())
}

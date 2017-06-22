package main

import (
	"fmt"
)

type person struct {
	Name    string
	Age     int
	Contact struct {
		Phone, City string
	}
}

func main() {
	a := person{Name: "yuxuefeng", Age: 33}
	a.Contact.Phone = "15901743422"
	a.Contact.City = "shanghai"
	fmt.Println(a)
}

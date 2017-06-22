package main

import (
	"fmt"
)

type USB interface {
	Name() string
	Connect()
}

type PhoneConnecter struct {
	name string
}

type TVConnecter struct{
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connet:", pc.name)
}

type (tv TVConnect)Connect(){
	fmt.Println("Connected:",tv.name)
}

func main() {
	var a USB
	a = PhoneConnecter{"PhoneConnecter"}
	a.Connect()
}

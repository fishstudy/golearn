package point

import "fmt"

func TestPoint() {
	var count int = 20
	var countPoint *int
	countPoint = &count
	fmt.Printf("count 的地址%x \n", &count)
	fmt.Printf("countPoint 的地址%x \n", countPoint)
	fmt.Printf("countPoint 的值%d \n", *countPoint)
}

func TestPointArr() {
	a, b := 1, 2
	pointArr := []*int{&a, &b}
	fmt.Println("指针数组 pointArr:", pointArr)
	arr := [...]int{3, 4, 5}
	arryPoint := &arr
	fmt.Println("数组指针 arrPoint", arryPoint)
}

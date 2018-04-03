package main

import (
	. "fmt"
)

type Class struct {
	strClassname string
	nClassNum    int
	strAdress    string
	next         *Class
}


func main() {

	sC := Class{"one", 200, "beijing", nil}
	Println(sC)

	sC1 := Class{
		strClassname: "two",
		nClassNum:    300,
	}
	sC2 := Class{
		strAdress:    "shanghai",
		strClassname: "three",
		nClassNum:    300,
		next:         &sC}
	Println(sC1, sC2)

//内嵌结构
	type Person struct {
		name string
		old  int
		XUHH struct {
			xu_name string
			xu_old  int
		}
	}

	p := Person{
		name: "luoxch",
		old:  1,
	}
	p.XUHH.xu_name = "nike"
	p.XUHH.xu_old = 2
	Println(p)

	var xu = struct {
		xu_name string
		xu_old  int
	}{
		xu_name: "Nike.Xu",
		xu_old:  200,
	}

//具有相同字段序列（字段名，类型，标签，顺序）的匿名struct 属于同一类型
	p.XUHH = xu // 同一类型的struct 可以进行 == ||  != ||  = 操作
	Println(p)

	//匿名字段
	type Resourse struct {
		id  string
		old int
	}
	type Classity struct {
		id string
	}
	type User struct {
		Resourse
		Classity
		name string
	}

	u := User{
		Resourse: Resourse{"10000", 1000},
		name:     "John",
	}
	// id 属于两个 struct 都共同拥有，所以这里不能直接u.id 需要加上struct名字
	//	u.Classity.id
	Print(u.Classity.id, "  ", u.Resourse.id, "  ", u.name, "  ", u.old, "  ", u.Resourse.id, "\r\n")

}

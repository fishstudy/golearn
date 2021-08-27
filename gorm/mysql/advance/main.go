package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID  int64
	Name sql.NullString `gorm:"default:'fish'"`
	Age int64
}
type Userbody struct {
	Name string
	Age int64
}
type Order struct {
	ID int64
	Uid int64
	Name string
}


func main() {
	//dsn := "root:root@tcp(127.0.0.1:3306)/go_learn?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_learn?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Order{});
	//order := Order{ID:1,Uid:2,Name:"bug"}
	//db.Table("orders").Create(&order)
	//fmt.Printf("%v",order)
	user := new(User)
	db.Debug().Where("ID=?",db.Table("orders").Select("ID").Where("Uid=?",2).SubQuery()).Find(user)
	fmt.Printf("%v",user)


    user2 := []User{}
    db.Debug().Table("users").Select("users.id,users.name,users.age").Joins("left join orders on users.id=orders.uid").Scan(&user2)
    fmt.Printf("%v",user2)
    fmt.Println("=====================================================")
	var names []string
    db.Table("users").Pluck("name",&names)
    fmt.Println(names)

	fmt.Println("=====================================================")
    var body []Userbody
	db.Debug().Table("users").Scan(&body)
    fmt.Printf("%v",body)
	fmt.Println("=======================---------------==============================")

    var body2 []Userbody
    db.Raw("select name,age from users ").Scan(&body2)
    fmt.Println(body2)
	fmt.Println("=====================================================")
    users := []User{}
    var count int64
    db.Debug().Table("users").Find(&users).Count(&count)
    fmt.Println(users)
    fmt.Println(count)

}

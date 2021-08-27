package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID  int64
	Name string
	Age int64
}

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_learn?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	var user User
	//db.Debug().First(&user,2)
	//user.Name="uuuu"
	//db.Debug().Save(&user)
	var user2 User
	db.Debug().Model(&user2).Where("id=?",1).Updates(map[string]interface{}{"name":"fish","age":36})
	fmt.Println(user2)
	var user3 User
	db.Debug().Model(&user3).Where("id=?",2).Updates(User{Name:"mama",Age:70})

	db.Debug().Delete(user)
}
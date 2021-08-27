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
	user := new(User)
	db.Debug().Where("Name=?", "fish").First(&user)
	fmt.Println(user)

	var user2 *User
	user2 = new(User)
	db.Debug().Where(&User{Name:sql.NullString{String:"fish",Valid:true}}).First(user2)

	user3 := new(User)
	db.Debug().Where(map[string]interface{}{"Name":"fish", "Age":18}).First(user3)

	user4 := []User{}
	db.Debug().Set("gorm:query_option", "FOR UPDATE").Where([]int{1,2,3}).Find(&user4)
	fmt.Printf("%v",user4)

	user5 := new(User)
	db.Debug().Attrs("Age",59).FirstOrInit(user5, &User{Name:sql.NullString{String:"fish",Valid:true}})
	fmt.Printf("%v",user5)

	user6 := new(User)
	db.Debug().Assign("Age",59).FirstOrInit(user6, &User{Name:sql.NullString{String:"fish",Valid:true}})
	fmt.Printf("%v",user6)
}

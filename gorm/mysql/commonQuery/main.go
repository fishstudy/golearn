package commonQuery

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

//func (u User)TableName() string {
//	return "person"
//}

func main() {
	//dsn := "root:root@tcp(127.0.0.1:3306)/go_learn?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_learn?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}
	//db.AutoMigrate(&User{})
	//user := User{Name:"fish",Age:18}
	//status := db.NewRecord(user)
	//fmt.Println(status)
	//db.Debug().Create(&user)
	//status = db.NewRecord(user)
	//fmt.Println(status)
	//user2 := User{Age:28}
	//db.Debug().Create(&user2)


	//user3 := User{Name:new(string),Age:28}
	//db.Debug().Create(&user3)

	//user4 := User{Name:sql.NullString{"",true},Age:28}
	//db.Debug().Create(&user4)

	//db.Table("man").CreateTable(&User{});
	//db.SingularTa(true)
	//db.AutoMigrate(&User{})
	user := new(User)
	user2 := new(User)
	user3 := new(User)
	user4 := new(User)
	user5 := new(User)
	db.Debug().First(&user)
	fmt.Printf("%v",user)
	fmt.Println("-------------first---------------------------------------------")

	db.Debug().Take(&user2)
	fmt.Printf("%v",user2)
	fmt.Println("----------------take------------------------------------------")

	db.Debug().Last(&user3)
	fmt.Printf("%v",user3)
	fmt.Println("--------------------last--------------------------------------")

	db.Debug().Find(&user4)
	fmt.Printf("%v",user4)
	fmt.Println("---------------------find-------------------------------------")

	db.Debug().First(&user5,3)
	fmt.Printf("%v",user5)
	fmt.Println("---------------------find---id----------------------------------")
}

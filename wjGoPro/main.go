package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type User struct {
	gorm.Model
	Name    string
	Age     int64
	Address string
}

func main() {
	//连接mysql
	dbType := "mysql"
	dsn := "root:wj0916@(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(dbType, dsn)
	if err != nil {
		fmt.Printf("连接数据库失败,err=%v\n", err)
		log.Fatalln("connect db error")
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("数据库关闭失败，err=%v", err)
			log.Fatalln("close db error")
		}
	}(db)
	db.AutoMigrate(&User{})
	user1 := User{Name: "wangjian", Age: 22, Address: "BeiJing"}
	result := db.Create(&user1)
	if result.Error != nil {
		fmt.Printf("create record error,err=%v", result.Error)
	}
}

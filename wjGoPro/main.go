package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

// PersonInfo 结构体对应数据表
type PersonInfo struct {
	ID      uint
	Name    string
	Address string
}

type User struct {
	gorm.Model   //内嵌gorm.Model
	Name         string
	Age          sql.NullInt64 //零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"` //指定类型;唯一索引(不能重复)
	Role         string  `gorm:"size:255"`                       //字段大小255
	MemberNumber *string `gorm:"unique;not null"`                //会员号唯一且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`                 //自增
	Address      string  `gorm:"index:addr"`                     //给address创建名字为addr的索引
	IgnoreMe     int     `gorm:"-"`                              //忽略本字段,不会映射到数据库表中
}

type Device struct {
	ID   int
	PID  int `gorm:"primary_key"`
	Name string
}

func (d Device) TableName() string {
	return "rice"
}

func main() {
	dbType := "mysql"
	dsn := "root:wj0916@(localhost:3306)/db_for_vvj?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(dbType, dsn)
	if err != nil {
		fmt.Printf("连接数据库失败,err=%v\n", err)
		log.Fatalln("connect db error")
	}
	db.SingularTable(true)
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("数据库关闭失败")
			log.Fatalln("db close error")
		}
	}(db)

	////创建表,利用AutoMigrate方法将结构体和数据库表进行对应
	//db.AutoMigrate(&PersonInfo{})
	////创建表行，将结构体实例和数据库行记录进行对应
	//p1 := PersonInfo{1, "王健", "北京市海淀区"}
	//db.Create(p1)
	//
	////查询表行
	//var p PersonInfo
	//db.First(&p)
	//fmt.Printf("p:%v\n", p)
	////更新
	//db.Model(&p).Update("address", "山东省烟台市")
	////删除
	//db.Delete(&p)
	//model := gorm.Model{}
	//println(model)
	//db.AutoMigrate(&User{})
	//db.AutoMigrate(&Device{})
	//db.Table("lol").CreateTable(&Device{})
	db.AutoMigrate(&PersonInfo{})
	p := PersonInfo{1, "王健", "北京市海淀区"}
	//打印p的主键是否已经在数据表中存在
	fmt.Println(db.NewRecord(p))
	db.Create(&p)
	fmt.Println(db.NewRecord(p))
	q := PersonInfo{1, "张良", "山东省烟台市"}
	fmt.Println(db.NewRecord(q))
	db.Create(&q)
}

package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type User struct {
	Id   int    `grom:"column:id"`
	Name string `gorm:"column:name"`
}

func main() {
	var path string = "root:password@tcp(writer:3306)/sample?charset=utf8&parseTime=true"
	db, err := gorm.Open(mysql.Open(path), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	_ = db

}

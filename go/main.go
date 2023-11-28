package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

type User struct {
	Id   int    `grom:"column:id"`
	Name string `gorm:"column:name"`
}

func main() {
	var writer string = "root:password@tcp(writer:3306)/sample?charset=utf8&parseTime=true"
	var reader string = "root:password@tcp(reader:3306)/sample?charset=utf8&parseTime=true"
	db, err := gorm.Open(mysql.Open(writer), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				ParameterizedQueries: true,
			},
		),
	})
	db.Logger = db.Logger.LogMode(logger.Info)
	db.Use(dbresolver.Register(dbresolver.Config{
		Replicas:          []gorm.Dialector{mysql.Open(reader)},
		Policy:            dbresolver.RandomPolicy{},
		TraceResolverMode: true,
	}))
	if err != nil {
		panic(err)
	}
	// res := db.Create(&User{Name: "test111"})
	// fmt.Println(res)
	// var Users []User
	// res = db.Model(&User{}).Find(&Users)
	// fmt.Printf("%+v", Users)
	var Users []User
	tx := db.Begin()
	tx.Create(&User{Name: "test111"})
	tx.Model(&User{Id: 1}).First(&Users)
	fmt.Printf("%+v", Users)
	tx.Commit()
}

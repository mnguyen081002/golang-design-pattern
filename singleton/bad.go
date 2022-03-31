package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type user struct {
	gorm.Model
	Name string
}

func getDBConnection() *gorm.DB {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&user{})
	return db
}

func main() {
	db := getDBConnection()
	db2 := getDBConnection()

	if db == db2 {
		fmt.Println("same")
	} else {
		fmt.Println("different")
	}
}

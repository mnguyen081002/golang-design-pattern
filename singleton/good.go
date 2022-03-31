package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sync"
)

var (
	once sync.Once
	db   *gorm.DB
)

type user2 struct {
	gorm.Model
	Name string
}

func getDBConnection2() *gorm.DB {
	once.Do(func() {
		db, _ = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		db.AutoMigrate(&user2{})
	})
	return db
}

func main() {
	db1 := getDBConnection2()
	db2 := getDBConnection2()

	if db1 == db2 {
		fmt.Println("same")
	} else {
		fmt.Println("different")
	}
}

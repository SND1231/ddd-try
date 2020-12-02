package main

import (
	"fmt"
	"ddd/v1/db"
	"ddd/v1/entity"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db := db.Connection()
	defer db.Close()
	fmt.Printf("migration")

	db.AutoMigrate(&model.User{})
}
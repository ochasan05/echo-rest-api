package main

import (
	"fmt"
	"go-rest-api-todo/db"
	"go-rest-api-todo/model"
)

func main() {
	defer fmt.Println("Successfully Migrated")

	dbConn := db.NewDB()
	defer db.CloseDB(dbConn)

	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Nigrated")
	defer db.Close(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}

package main

import (
	"fmt"
	"log"

	"github.com/VitorC-sa/crud-meplux/database"
	"github.com/VitorC-sa/crud-meplux/routers"
)

func connectDB() {
	if err := database.ConnectDB(); err != nil {
		msg := fmt.Sprintf("Cannot connect to Database - %s", err.Error())
		log.Fatal(msg)
	}
}

func main() {
	connectDB()
	routers.HandleRoutes()
}

package main

import (
	"github.com/ssjlee93/fitworks-data-user/configs"
	"log"
)

func main() {
	// Code
	log.Println("Starting the application...")
	db := configs.GetConnection()
	log.Println(db.Stats().OpenConnections)
}

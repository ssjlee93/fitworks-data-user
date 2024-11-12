package main

import (
	"fmt"
	"github.com/ssjlee93/fitworks-data-user/configs"
	"github.com/ssjlee93/fitworks-data-user/daos"
	"log"
)

func main() {
	// Code
	log.Println("Starting the application...")
	db := configs.GetConnection()
	log.Println(db.Stats().OpenConnections)

	roleDao := daos.NewRoleDAOImpl(db)

	roles, err := roleDao.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found %d roles\n", len(roles))
}

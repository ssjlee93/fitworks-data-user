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

	roles, err := roleDao.ReadOne(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found %q roles\n", roles.Role)
}

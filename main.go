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
	defer db.Close()

	// placeholder for var db
	fmt.Println(db.Stats().OpenConnections)

	roleDao := daos.NewRoleDAOImpl(db)
	userDao := daos.NewUserDAOImpl(db)

	// placeholder for var roleDao
	res, err := roleDao.ReadAll()
	if err != nil {
		errorMsg := fmt.Errorf("main error ReadAll: %v", err)
		fmt.Println(errorMsg)
	}
	for _, role := range res {
		role.PrintRole()
	}

	users, err := userDao.ReadAll()
	if err != nil {
		errorMsg := fmt.Errorf("main error ReadAll: %v", err)
		fmt.Println(errorMsg)
	}
	for _, user := range users {
		user.PrintUser()
	}

}

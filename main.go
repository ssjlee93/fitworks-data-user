package main

import (
	"log"
	"net/http"

	"github.com/ssjlee93/fitworks-data-user/configs"
	"github.com/ssjlee93/fitworks-data-user/controllers"
	"github.com/ssjlee93/fitworks-data-user/daos"
	"github.com/ssjlee93/fitworks-data-user/repositories"
)

func main() {
	// Code
	log.Println("Starting the application...")
	db := configs.GetConnection()
	defer db.Close()

	userDao := daos.NewUserDAOImpl(db)

	userRepo := repositories.NewUserRepository(*userDao)

	userController := controllers.NewUserController(*userRepo)

	http.HandleFunc("/users", userController.ReadAllHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

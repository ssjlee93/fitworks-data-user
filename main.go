package main

import (
	"log"
	"net/http"

	"github.com/ssjlee93/fitworks-data-user/configs"
	"github.com/ssjlee93/fitworks-data-user/controllers"
	"github.com/ssjlee93/fitworks-data-user/repositories"
	"github.com/ssjlee93/fitworks-data-user/services"
)

func main() {
	// Code
	log.Println("Starting the application...")
	db := configs.GetConnection()
	defer db.Close()

	userDao := repositories.NewUserDAOImpl(db)

	userRepo := services.NewUserRepository(*userDao)

	userController := controllers.NewUserController(*userRepo)

	http.HandleFunc("/users", userController.ReadAllHandler)
	http.HandleFunc("/user/", userController.Handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

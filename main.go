package main

import (
	"ProjectManagement/config"
	"ProjectManagement/controllers"
	"ProjectManagement/database/seed"
	"ProjectManagement/repositories"
	"ProjectManagement/routes"
	"ProjectManagement/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config.LoadEnv()
	config.ConnectDB()

	seed.SeedAdmin()
	app := fiber.New()

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	routes.Setup(app, userController)
	port := config.AppConfig.AppPort
	log.Println("server is running at :", port)
	log.Fatal(app.Listen(":" + port))
}

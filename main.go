package main

import (
	"ProjectManagement/config"
	"ProjectManagement/controllers"
	"ProjectManagement/repositories"
	"ProjectManagement/routes"
	"ProjectManagement/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config.LoadEnv()
	config.ConnectDB()

	app := fiber.New()

	// User
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// board
	boardRepo := repositories.NewBoardRepository()
	boardMemberRepo := repositories.NewBoardMemberRepository()
	boardService := services.NewBoardService(boardRepo, userRepo, boardMemberRepo)
	boardController := controllers.NewBoardController(boardService)

	routes.Setup(app, userController, boardController)
	port := config.AppConfig.AppPort
	log.Println("server is running at :", port)
	log.Fatal(app.Listen(":" + port))
}

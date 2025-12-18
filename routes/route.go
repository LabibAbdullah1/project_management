package routes

import (
	"ProjectManagement/config"
	"ProjectManagement/controllers"
	"ProjectManagement/utils"
	"log"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
)

func Setup(app *fiber.App, uc *controllers.UserController) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env failed")
	}
	app.Post("/v1/auth/register", uc.Register)
	app.Post("/v1/auth/login", uc.Login)

	//JWT Protected Routes
	api := app.Group("/api/v1", jwtware.New(jwtware.Config{
		SigningKey: []byte(config.AppConfig.JWTSecret),
		ContextKey: "user",
		ErrorHandler: func(c *fiber.Ctx, err error) error{
			return utils.Unauthorized(c,"Error Unauthorized", err.Error())
		},
	}))

	userGroup := api.Group("/user")
	userGroup.Get("/:id", uc.GetUser)
}

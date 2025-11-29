package seed

import (
	"ProjectManagement/config"
	"ProjectManagement/models"
	"ProjectManagement/utils"
	"log"
)

func SeedAdmin() {
	password, _ := utils.HashPassword("password123")

	admin := models.User{
		Name:     "Admin User",
		Email:    "admin@example.com",
		Password: password,
		Role:     "admin",
	}
	if err := config.DB.FirstOrCreate(&admin, models.User{Email: admin.Email}); 
	err != nil {
		log.Println("Failed to seed admin", err)
	} else {
		log.Println("Admin seeded successfully")
	}
}

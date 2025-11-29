package utils

import (
	"ProjectManagement/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

//generate JWT token
//generate Refresh Token

func GenerateToken(userID int64, role, email string, publicID uuid.UUID) {
	secret := config.AppConfig.JWTSecret
	duration, _ := time.ParseDuration(config.AppConfig.JWTExpire)

	claims := JWT.MapClaims{
		"user_id": userID,
		"role":    role,
		"email":   email,
		"pub_id":  publicID,
		"exp":     time.Now().Add(duration).Unix(),
	}

	token := JWT.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))

}

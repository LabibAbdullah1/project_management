package models

import (
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type User struct {
	InternalID int64          `json:"internal_id" db:"internal_id" gorm:"PrimaryKey"`
	PublicID   uuid.UUID      `json:"public_id" db:"public_id"`
	Name       string         `json:"name" db:"name"`
	Email      string         `json:"email" db:"email"`
	Password   string         `json:"password" db:"password" gorm:"column:password"`
	Role       string         `json:"role" db:"role"`
	CreatedAt  time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at" db:"updated-at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}

type UseResponse struct {
	PublicID  uuid.UUID      `json:"public_id" `
	Name      string         `json:"name" `
	Email     string         `json:"email" `
	Role      string         `json:"role" `
	CreatedAt time.Time      `json:"created_at" `
	UpdatedAt time.Time      `json:"updated_at" `
	DeletedAt gorm.DeletedAt `json:"-"`
}

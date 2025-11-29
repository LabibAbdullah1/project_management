package models

import (
	"github.com/google/uuid"
	"time"
)

type Card struct {
	InternalID int64 `json:"InternalID" db:"internal_id" gorm:"primaryKey;autoIncrement"`
	PublicID   uuid.UUID `json:"PublicID" db:"public_id"`
	ListID    int64 `json:"list-internal_id" db:"list_internal_id" gorm:"column:list_internal_id;not null;index"`
	Title     string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	DueDate   *time.Time `json:"due_date,omitempty" db:"due_date"`
	Position  int `json:"position" db:"position"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
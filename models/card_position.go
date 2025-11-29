package models

import (
	"ProjectManagement/models/types"

	"github.com/google/uuid"
)

type CardPosition struct {
	InternalID int64 `json:"internal_id" db:"internal_id" gorm:"primaryKey;autoIncrement"`
	PublicID   uuid.UUID `json:"public_id" db:"public_id" gorm:"type:uuid;not null"`
	ListID    int64 `json:"list_id" db:"list_id" gorm:"column:list_internal_id;not null;not null"`
	CardOrder types.UUIDArray `json:"card_order" db:"type:uuid[]"`
}
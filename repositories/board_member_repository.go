package repositories

import (
	"ProjectManagement/config"
	"ProjectManagement/models"
)

type BoardMemberRepository interface {
	GetMember(boardPublicID string) ([]models.User, error)
}

type boardMemberRepository struct {
}

func NewBoardMemberRepository() BoardMemberRepository {
	return &boardMemberRepository{}
}

func (r *boardMemberRepository) GetMember(boardPublicID string) ([]models.User, error) {
	var users []models.User
	err := config.DB.Model(&models.User{}).
		Joins("JOIN board_members ON board_members.user_internal_id = users.internal_id").
		Joins("JOIN boards ON boards.internal_id = board_members.board_internal_id").
		Where("boards.public_id = ?", boardPublicID).
		Find(&users).Error
	return users, err

}

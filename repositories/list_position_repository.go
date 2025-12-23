package repositories

import (
	"ProjectManagement/config"
	"ProjectManagement/models"

	"github.com/google/uuid"
)

type listPositionRepository struct {
}

type ListPositionRepository interface {
	GetByBoard(boardPublicID string) (*models.ListPosition, error)
	CreateOrUpdate(boardPublicID string, listorder []uuid.UUID) error
	GetListOrder(boardPublicID string) ([]uuid.UUID, error)
	UpdateListOrder(position *models.ListPosition) error
}

func NewListPositionrepository() ListPositionRepository {
	return &listPositionRepository{}
}

func (r *listPositionRepository) GetByBoard(boardPublicID string) (*models.ListPosition, error) {
	var position models.ListPosition

	err := config.DB.Joins("JOIN board ON boards.internal_id = list_positions.board_internal_id").
		Where("boards.public_id = ?", boardPublicID).Error
	return &position, err
}

func (r *listPositionRepository) CreateOrUpdate(boardPublicID string, listorder []uuid.UUID) error {
	return config.DB.Exec(`
	INSERT INTO list_position (board_internal_id , list_order)
	SELECT internal_id, ? FROM board where public_id = ? 
	ON CONFLICT (board_internal_id) 
	DO UPDATE SET list-order = EXCLUDE.list_order`, listorder, boardPublicID).Error
}

func (r *listPositionRepository) GetListOrder(boardPublicID string) ([]uuid.UUID, error) {
	position, err := r.GetByBoard(boardPublicID)
	if err != nil {
		return nil, err
	}
	return position.ListOrder, err
}

func (r *listPositionRepository) UpdateListOrder(position *models.ListPosition) error {
	return config.DB.Model(position).
		Where("internal_id = ?", position.InternalID).
		Update("list_order", position.ListOrder).Error
}

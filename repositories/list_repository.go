package repositories

import (
	"ProjectManagement/config"
	"ProjectManagement/models"
)

type ListRepository interface {
}
type listRepository struct {
}

func NewListRepository() ListRepository {
	return &listRepository{}
}

func (r *listRepository) Create(list *models.List) error {
	return config.DB.Create(list).Error
}

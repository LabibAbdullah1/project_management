package services

import (
	"ProjectManagement/models"
	"ProjectManagement/repositories"

	"github.com/google/uuid"
)

type ListService interface {
	GetByBoardID(board_public_id string) (*listWithOrder, error)
	GetByID(id uint) (*models.List, error)
	GetByPublicID(publicID string) (*models.List, error)
	Create(list *models.List) error
	Update(list *models.List) error
	Delete(id uint) error
	UpdatePosition(boardPublicID string, position []uuid.UUID) error
}

type listService struct {
	listRepo    repositories.ListRepository
	boardRepo   repositories.BoardRepository
	listPosRepo repositories.ListPositionRepository
}

func NewListService(
	listRepo repositories.ListRepository,
	boardRepo repositories.BoardRepository,
	listPosRepo repositories.ListPositionRepository,
) ListService {
	return &listService{ listRepo, boardRepo, listPosRepo }
}

type listWithOrder struct {
	Position []uuid.UUID
	Lists    []models.List
}

func (s *listService) GetByBoardID(board_public_id string) (*listWithOrder, error) {
	// TODO: implement GetByBoardID
	return nil, nil
}

func (s *listService) GetByID(id uint) (*models.List, error) {
	// TODO: implement GetByID
	return nil, nil
}

func (s *listService) GetByPublicID(publicID string) (*models.List, error) {
	// TODO: implement GetByPublicID
	return nil, nil
}

func (s *listService) Create(list *models.List) error {
	// TODO: implement Create
	return nil
}

func (s *listService) Update(list *models.List) error {
	// TODO: implement Update
	return nil
}

func (s *listService) Delete(id uint) error {
	// TODO: implement Delete
	return nil
}

func (s *listService) UpdatePosition(boardPublicID string, position []uuid.UUID) error {
	// TODO: implement UpdatePosition
	return nil
}

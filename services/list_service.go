package services

import (
	"ProjectManagement/config"
	"ProjectManagement/models"
	"ProjectManagement/models/types"
	"ProjectManagement/repositories"
	"ProjectManagement/utils"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ListService interface {
	GetByBoardID(board_public_id string) (*ListWithOrder, error)
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
	return &listService{listRepo, boardRepo, listPosRepo}
}

type ListWithOrder struct {
	Position []uuid.UUID
	Lists    []models.List
}

func (s *listService) GetByBoardID(boardPublicID string) (*ListWithOrder, error) {
	// TODO: implement GetByBoardID
	_, err := s.boardRepo.FindByPublicID(boardPublicID)
	if err != nil {
		return nil, errors.New("Board Not Found")
	}

	position, err := s.listPosRepo.GetListOrder(boardPublicID)
	if err != nil {
		return nil, errors.New("Failed To Get list order :" + err.Error())
	}

	lists, err := s.listRepo.FindByBoardID(boardPublicID)
	if err != nil {
		return nil, errors.New("Failed To Get List :" + err.Error())
	}

	// sorting by position
	orderedList := utils.SortingListByPosition(lists, position)

	return &ListWithOrder{
		Position: position,
		Lists:    orderedList,
	}, nil
}

func (s *listService) GetByID(id uint) (*models.List, error) {
	return s.listRepo.FindByID(id)
}

func (s *listService) GetByPublicID(publicID string) (*models.List, error) {
	return s.listRepo.FindByPublicID(publicID)
}

func (s *listService) Create(list *models.List) error {
	// validasi board
	board, err := s.boardRepo.FindByPublicID(list.BoardPublicID.String())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("Board not Found ")
		}
		return fmt.Errorf("Failed to Get Board :%w", err)
	}
	list.BoardInternalID = board.InternalID

	if list.PublicID == uuid.Nil {
		list.PublicID = uuid.New()
	}

	tx := config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// simpan list baru
	if err := tx.Create(list).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("Failed To Create Lsit : %w", err)
	}

	// update position
	var position models.ListPosition

	res := tx.Where("board_internal_id = ?", board.InternalID).First(&position)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		// buatbaru jika belum ada
		position = models.ListPosition{
			PublicID:  uuid.New(),
			BoardID:   board.InternalID,
			ListOrder: types.UUIDArray{list.PublicID},
		}
		if err := tx.Create(&position).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("Failed to create list position : %w", err)
		}
	} else if res.Error != nil {
		tx.Rollback()
		return fmt.Errorf("Failed to create list position : %w", res.Error)
	} else {
		// tambahkan ID baru
		position.ListOrder = append(position.ListOrder, list.PublicID)
		if err := tx.Model(&position).Update("list_order", position.ListOrder).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("Failed to update list position : %w", err)
		}
	}

	// commit transaction
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("Transaction Commit Failed %w", err)
	}

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

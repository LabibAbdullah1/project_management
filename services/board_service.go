package services

import (
	"ProjectManagement/models"
	"ProjectManagement/repositories"
	"errors"

	"github.com/google/uuid"
)

type BoardService interface {
	Create(board *models.Board) error
	Update(board *models.Board) error
	GetByPublicID(publicID string)(*models.Board, error)
	AddMember(boardPublicID string, userPublicIDS []string) error
	RemoveMembers(boardPublicID string, userPublicIDs [] string) error
}

type boardService struct {
	boardRepo repositories.BoardRepository
	userRepo  repositories.UserRepository
	boardMemberRepo repositories.BoardMemberRepository
}

func NewBoardService(
	boardRepo repositories.BoardRepository, 
	userRepo repositories.UserRepository,
	boardMemberRepo repositories.BoardMemberRepository,
	) BoardService {
	return &boardService{boardRepo, userRepo, boardMemberRepo}
}
func (s *boardService) Create(board *models.Board) error {
	user, err := s.userRepo.FindByPublicID(board.OwnerPublicID.String())
	if err != nil {
		return errors.New("Owner not Found")
	}
	board.PublicID = uuid.New()
	board.OwnerID = user.InternalID
	return s.boardRepo.Create(board)
}

func (s *boardService) Update(board *models.Board) error {
	return s.boardRepo.Update(board)
}
func (s *boardService) GetByPublicID(publicID string)(*models.Board, error){
	return s.boardRepo.FindByPublicID(publicID)
}
func (s *boardService) AddMember(boardPublicID string, userPublicIDS []string) error {
	board, err := s.boardRepo.FindByPublicID(boardPublicID)
	if err != nil {
		return errors.New("Board is Not Found")
	}

	var userInternalIDs []uint
	for _, userPublicID := range userPublicIDS{
		user, err := s.userRepo.FindByPublicID(userPublicID)
		if err != nil {
			return errors.New("User Not Found" + userPublicID)
		}
		userInternalIDs = append(userInternalIDs, uint(user.InternalID))
	}
	// cek ke anggotaan
	existingMember, err :=	s.boardMemberRepo.GetMember(string(board.PublicID.String()))
	if err != nil {
		return err
	}

	// cek menggunakan map
	memberMap := make(map[uint]bool)
	for _, member := range existingMember{
		memberMap[uint(member.InternalID)] = true
	}

	var newMemberIDs []uint
	for _,userID := range userInternalIDs {
		if !memberMap[userID]{
			newMemberIDs = append(newMemberIDs, userID)
		}
	}
	if len (newMemberIDs) == 0 {
		return nil
	}
	return s.boardRepo.AddMember(uint(board.InternalID), newMemberIDs)
}

func (s *boardService) RemoveMembers(boardPublicID string, userPublicIDs [] string) error {
	board, err := s.boardRepo.FindByPublicID(boardPublicID)
	if err != nil {
		return errors.New("Board is Not Found")
	}

	//validasi user
	var userInternalIDs []uint
	for _, userPublicID := range userPublicIDs{
		user, err := s.userRepo.FindByPublicID(userPublicID)
		if err != nil {
			return errors.New("User Not Found" + userPublicID)
		}
		userInternalIDs = append(userInternalIDs, uint(user.InternalID))
	}

	//cek keanggotaan
	existingMember, err :=	s.boardMemberRepo.GetMember(string(board.PublicID.String()))
	if err != nil {
		return err
	}

	// cek menggunakan map
	memberMap := make(map[uint]bool)
	for _, member := range existingMember{
		memberMap[uint(member.InternalID)] = true
	}

	var membersToRemove []uint
	for _, userID := range userInternalIDs {
		if memberMap[userID] {
			membersToRemove = append(membersToRemove, userID)
		}
	}
	return s.boardRepo.RemoveMembers(uint(board.InternalID), membersToRemove)
}
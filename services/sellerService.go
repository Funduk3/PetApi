package services

import (
	"errors"

	"petstore-api/models"
	"petstore-api/repositories"

	"gorm.io/gorm"
)

type sellerService struct {
	sellerRepo repositories.UserRepository
	petRepo    repositories.PetRepository
}

func NewSellerService(sellerRepo repositories.UserRepository, petRepo repositories.PetRepository) UserService {
	return &sellerService{
		sellerRepo: sellerRepo,
		petRepo:    petRepo,
	}
}

func (s *sellerService) GetAll(includePets bool) ([]models.User, error) {
	return s.sellerRepo.GetAll(includePets)
}

func (s *sellerService) GetByID(id uint, includePets bool) (*models.User, error) {
	seller, err := s.sellerRepo.GetByID(id, includePets)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("seller not found")
		}
		return nil, err
	}
	return seller, nil
}

func (s *sellerService) Create(req *models.CreateUserRequest) (*models.User, error) {
	if req.Name == "" || req.Email == "" {
		return nil, errors.New("name and email are required")
	}

	seller := &models.User{
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Address: req.Address,
	}

	err := s.sellerRepo.Create(seller)
	if err != nil {
		return nil, err
	}

	return seller, nil
}

func (s *sellerService) Update(id uint, req *models.UpdateUserRequest) (*models.User, error) {
	seller, err := s.sellerRepo.GetByID(id, false)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("seller not found")
		}
		return nil, err
	}

	if req.Name != "" {
		seller.Name = req.Name
	}
	if req.Email != "" {
		seller.Email = req.Email
	}
	if req.Phone != "" {
		seller.Phone = req.Phone
	}
	if req.Address != "" {
		seller.Address = req.Address
	}

	err = s.sellerRepo.Update(seller)
	if err != nil {
		return nil, err
	}

	return seller, nil
}

func (s *sellerService) Delete(id uint) error {
	_, err := s.sellerRepo.GetByID(id, false)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("seller not found")
		}
		return err
	}

	return s.sellerRepo.Delete(id)
}

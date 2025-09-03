package services

import (
	"errors"

	"petstore-api/models"
	"petstore-api/repositories"

	"gorm.io/gorm"
)

type buyerService struct {
	buyerRepo repositories.UserRepository
}

func (b *buyerService) GetAll(includePets bool) ([]models.User, error) {
	return b.buyerRepo.GetAll(false)
}

func (b *buyerService) GetByID(id uint, includePets bool) (*models.User, error) {
	buyer, err := b.buyerRepo.GetByID(id, false)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("seller not found")
		}
		return nil, err
	}
	return buyer, nil
}

func (b *buyerService) Create(req *models.CreateUserRequest) (*models.User, error) {
	if req.Name == "" || req.Email == "" {
		return nil, errors.New("name and email are required")
	}

	buyer := &models.User{
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Address: req.Address,
	}

	err := b.buyerRepo.Create(buyer)
	if err != nil {
		return nil, err
	}

	return buyer, nil
}

func (b *buyerService) Update(id uint, req *models.UpdateUserRequest) (*models.User, error) {
	buyer, err := b.buyerRepo.GetByID(id, false)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("buyer not found")
		}
		return nil, err
	}

	if req.Name != "" {
		buyer.Name = req.Name
	}
	if req.Email != "" {
		buyer.Email = req.Email
	}
	if req.Phone != "" {
		buyer.Phone = req.Phone
	}
	if req.Address != "" {
		buyer.Address = req.Address
	}

	err = b.buyerRepo.Update(buyer)
	if err != nil {
		return nil, err
	}

	return buyer, nil
}

func (b *buyerService) Delete(id uint) error {
	_, err := b.buyerRepo.GetByID(id, false)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("seller not found")
		}
		return err
	}
	return b.buyerRepo.Delete(id)
}

func NewBuyerService(buyerRepo repositories.UserRepository) UserService {
	return &buyerService{
		buyerRepo: buyerRepo,
	}
}

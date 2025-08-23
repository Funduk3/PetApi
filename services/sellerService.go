package services

import (
	"errors"

	"petstore-api/models"
	"petstore-api/repositories"

	"gorm.io/gorm"
)

type sellerService struct {
	sellerRepo repositories.SellerRepository
	petRepo    repositories.PetRepository
}

func NewSellerService(sellerRepo repositories.SellerRepository, petRepo repositories.PetRepository) SellerService {
	return &sellerService{
		sellerRepo: sellerRepo,
		petRepo:    petRepo,
	}
}

func (s *sellerService) GetAllSellers(includePets bool) ([]models.Seller, error) {
	return s.sellerRepo.GetAll(includePets)
}

func (s *sellerService) GetSellerByID(id uint, includePets bool) (*models.Seller, error) {
	seller, err := s.sellerRepo.GetByID(id, includePets)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("seller not found")
		}
		return nil, err
	}
	return seller, nil
}

func (s *sellerService) CreateSeller(req *models.CreateSellerRequest) (*models.Seller, error) {
	if req.Name == "" || req.Email == "" {
		return nil, errors.New("name and email are required")
	}

	seller := &models.Seller{
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

func (s *sellerService) UpdateSeller(id uint, req *models.UpdateSellerRequest) (*models.Seller, error) {
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

func (s *sellerService) DeleteSeller(id uint) error {
	_, err := s.sellerRepo.GetByID(id, false)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("seller not found")
		}
		return err
	}

	petCount, err := s.sellerRepo.GetPetCount(id)
	if err != nil {
		return err
	}

	if petCount > 0 {
		return errors.New("cannot delete seller with existing pets")
	}

	return s.sellerRepo.Delete(id)
}

package services

import (
	"errors"

	"petstore-api/models"
	"petstore-api/repositories"

	"gorm.io/gorm"
)

type petService struct {
	petRepo    repositories.PetRepository
	sellerRepo repositories.UserRepository
}

func NewPetService(petRepo repositories.PetRepository, sellerRepo repositories.UserRepository) PetService {
	return &petService{
		petRepo:    petRepo,
		sellerRepo: sellerRepo,
	}
}

func (s *petService) GetAllPets(includeSeller bool, sellerID *uint) ([]models.Pet, error) {
	return s.petRepo.GetAll(includeSeller, sellerID)
}

func (s *petService) GetPetByID(id uint, includeSeller bool) (*models.Pet, error) {
	pet, err := s.petRepo.GetByID(id, includeSeller)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("pet not found")
		}
		return nil, err
	}
	return pet, nil
}

func (s *petService) CreatePet(req *models.CreatePetRequest) (*models.Pet, error) {
	if req.Name == "" || req.Species == "" || req.SellerID == 0 {
		return nil, errors.New("name, species, and seller_id are required")
	}

	_, err := s.sellerRepo.GetByID(req.SellerID, false)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("seller not found")
		}
		return nil, err
	}

	pet := &models.Pet{
		Name:        req.Name,
		Species:     req.Species,
		Breed:       req.Breed,
		Age:         req.Age,
		Price:       req.Price,
		Description: req.Description,
		Available:   req.Available,
		SellerID:    req.SellerID,
	}

	err = s.petRepo.Create(pet)
	if err != nil {
		return nil, err
	}

	return pet, nil
}

func (s *petService) UpdatePet(id uint, req *models.UpdatePetRequest) (*models.Pet, error) {
	pet, err := s.petRepo.GetByID(id, false)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("pet not found")
		}
		return nil, err
	}

	if req.Name != "" {
		pet.Name = req.Name
	}
	if req.Species != "" {
		pet.Species = req.Species
	}
	if req.Breed != "" {
		pet.Breed = req.Breed
	}
	if req.Age > 0 {
		pet.Age = req.Age
	}
	if req.Price > 0 {
		pet.Price = req.Price
	}
	if req.Description != "" {
		pet.Description = req.Description
	}
	if req.Available != nil {
		pet.Available = *req.Available
	}

	if req.SellerID != 0 && req.SellerID != pet.SellerID {
		_, err := s.sellerRepo.GetByID(req.SellerID, false)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, errors.New("seller not found")
			}
			return nil, err
		}
		pet.SellerID = req.SellerID
	}

	err = s.petRepo.Update(pet)
	if err != nil {
		return nil, err
	}

	return pet, nil
}

func (s *petService) DeletePet(id uint) error {
	_, err := s.petRepo.GetByID(id, false)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("pet not found")
		}
		return err
	}

	return s.petRepo.Delete(id)
}

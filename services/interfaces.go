package services

import "petstore-api/models"

type SellerService interface {
	GetAllSellers(includePets bool) ([]models.Seller, error)
	GetSellerByID(id uint, includePets bool) (*models.Seller, error)
	CreateSeller(req *models.CreateSellerRequest) (*models.Seller, error)
	UpdateSeller(id uint, req *models.UpdateSellerRequest) (*models.Seller, error)
	DeleteSeller(id uint) error
}

type PetService interface {
	GetAllPets(includeSeller bool, sellerID *uint) ([]models.Pet, error)
	GetPetByID(id uint, includeSeller bool) (*models.Pet, error)
	CreatePet(req *models.CreatePetRequest) (*models.Pet, error)
	UpdatePet(id uint, req *models.UpdatePetRequest) (*models.Pet, error)
	DeletePet(id uint) error
}

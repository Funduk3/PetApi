package repositories

import "petstore-api/models"

type SellerRepository interface {
	GetAll(includePets bool) ([]models.Seller, error)
	GetByID(id uint, includePets bool) (*models.Seller, error)
	Create(seller *models.Seller) error
	Update(seller *models.Seller) error
	Delete(id uint) error
	GetPetCount(sellerID uint) (int64, error)
}

type PetRepository interface {
	GetAll(includeSeller bool, sellerID *uint) ([]models.Pet, error)
	GetByID(id uint, includeSeller bool) (*models.Pet, error)
	Create(pet *models.Pet) error
	Update(pet *models.Pet) error
	Delete(id uint) error
	GetBySellerID(sellerID uint) ([]models.Pet, error)
}

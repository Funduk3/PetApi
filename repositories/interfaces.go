package repositories

import "petstore-api/models"

type UserRepository interface {
	GetAll(includePets bool) ([]models.User, error)
	GetByID(id uint, includePets bool) (*models.User, error)
	Create(seller *models.User) error
	Update(seller *models.User) error
	Delete(id uint) error
}

type PetRepository interface {
	GetAll(includeSeller bool, sellerID *uint) ([]models.Pet, error)
	GetByID(id uint, includeSeller bool) (*models.Pet, error)
	Create(pet *models.Pet) error
	Update(pet *models.Pet) error
	Delete(id uint) error
	GetBySellerID(sellerID uint) ([]models.Pet, error)
}

type UserItemRepository interface {
	AddPet(userID uint, petID uint) error
	RemovePet(userID uint, petID uint) error
	GetPetsByBuyerID(userID uint) ([]models.Pet, error)
}

package services

import "petstore-api/models"

type UserService interface {
	GetAll(includePets bool) ([]models.User, error)
	GetByID(id uint, includePets bool) (*models.User, error)
	Create(req *models.CreateUserRequest) (*models.User, error)
	Update(id uint, req *models.UpdateUserRequest) (*models.User, error)
	Delete(id uint) error
}

type PetService interface {
	GetAllPets(includeSeller bool, sellerID *uint) ([]models.Pet, error)
	GetPetByID(id uint, includeSeller bool) (*models.Pet, error)
	CreatePet(req *models.CreatePetRequest) (*models.Pet, error)
	UpdatePet(id uint, req *models.UpdatePetRequest) (*models.Pet, error)
	DeletePet(id uint) error
}

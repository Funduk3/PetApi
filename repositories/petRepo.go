package repositories

import (
	"gorm.io/gorm"
	"petstore-api/models"
)

type petRepository struct {
	db *gorm.DB
}

func NewPetRepository(db *gorm.DB) PetRepository {
	return &petRepository{db: db}
}

func (r *petRepository) GetAll(includeSeller bool, sellerID *uint) ([]models.Pet, error) {
	var pets []models.Pet
	query := r.db

	if includeSeller {
		query = query.Preload("Seller")
	}

	if sellerID != nil {
		query = query.Where("seller_id = ?", *sellerID)
	}

	result := query.Find(&pets)
	return pets, result.Error
}

func (r *petRepository) GetByID(id uint, includeSeller bool) (*models.Pet, error) {
	var pet models.Pet
	query := r.db

	if includeSeller {
		query = query.Preload("Seller")
	}

	result := query.First(&pet, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &pet, nil
}

func (r *petRepository) Create(pet *models.Pet) error {
	result := r.db.Create(pet)
	return result.Error
}

func (r *petRepository) Update(pet *models.Pet) error {
	result := r.db.Save(pet)
	return result.Error
}

func (r *petRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Pet{}, id)
	return result.Error
}

func (r *petRepository) GetBySellerID(sellerID uint) ([]models.Pet, error) {
	var pets []models.Pet
	result := r.db.Where("seller_id = ?", sellerID).Find(&pets)
	return pets, result.Error
}

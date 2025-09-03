package users

import (
	"gorm.io/gorm"
	"petstore-api/models"
	"petstore-api/repositories"
)

type sellerRepository struct {
	db *gorm.DB
}

func NewSellerRepository(db *gorm.DB) repositories.UserRepository {
	return &sellerRepository{db: db}
}

func (r *sellerRepository) GetAll(includePets bool) ([]models.User, error) {
	var sellers []models.User
	query := r.db

	if includePets {
		query = query.Preload("Pets")
	}

	result := query.Find(&sellers)
	return sellers, result.Error
}

func (r *sellerRepository) GetByID(id uint, includePets bool) (*models.User, error) {
	var seller models.User
	query := r.db

	if includePets {
		query = query.Preload("Pets")
	}

	result := query.First(&seller, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &seller, nil
}

func (r *sellerRepository) Create(seller *models.User) error {
	result := r.db.Create(seller)
	return result.Error
}

func (r *sellerRepository) Update(seller *models.User) error {
	result := r.db.Save(seller)
	return result.Error
}

func (r *sellerRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Seller{}, id)
	return result.Error
}

func (r *sellerRepository) GetPetCount(sellerID uint) (int64, error) {
	var count int64
	result := r.db.Model(&models.Pet{}).Where("seller_id = ?", sellerID).Count(&count)
	return count, result.Error
}

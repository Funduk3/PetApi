package repositories

import (
	"gorm.io/gorm"
	"petstore-api/models"
)

type sellerRepository struct {
	db *gorm.DB
}

func NewSellerRepository(db *gorm.DB) SellerRepository {
	return &sellerRepository{db: db}
}

func (r *sellerRepository) GetAll(includePets bool) ([]models.Seller, error) {
	var sellers []models.Seller
	query := r.db

	if includePets {
		query = query.Preload("Pets")
	}

	result := query.Find(&sellers)
	return sellers, result.Error
}

func (r *sellerRepository) GetByID(id uint, includePets bool) (*models.Seller, error) {
	var seller models.Seller
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

func (r *sellerRepository) Create(seller *models.Seller) error {
	result := r.db.Create(seller)
	return result.Error
}

func (r *sellerRepository) Update(seller *models.Seller) error {
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

package users

import (
	"gorm.io/gorm"
	"petstore-api/models"
	"petstore-api/repositories"
)

type buyerRepository struct {
	db *gorm.DB
}

func (b buyerRepository) GetAll(includePets bool) ([]models.User, error) {
	var buyers []models.User
	query := b.db

	if includePets {
		query = query.Preload("Pets")
	}
	result := query.Find(&buyers)
	return buyers, result.Error
}

func (b buyerRepository) GetByID(id uint, includePets bool) (*models.User, error) {
	var buyers models.User
	query := b.db

	if includePets {
		query = query.Preload("Pets")
	}

	result := query.First(&buyers, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &buyers, nil
}

func (b buyerRepository) Create(seller *models.User) error {
	result := b.db.Create(seller)
	return result.Error
}

func (b buyerRepository) Update(seller *models.User) error {
	result := b.db.Save(seller)
	return result.Error
}

func (b buyerRepository) Delete(id uint) error {
	result := b.db.Delete(&models.User{}, id)
	return result.Error
}

func NewBuyerRepository(db *gorm.DB) repositories.UserRepository {
	return &buyerRepository{db: db}
}

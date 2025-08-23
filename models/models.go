package models

import (
	"time"
)

type Seller struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"not null;size:255"`
	Email     string    `json:"email" gorm:"unique;not null;size:255"`
	Phone     string    `json:"phone" gorm:"size:20"`
	Address   string    `json:"address" gorm:"size:500"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Pets      []Pet     `json:"pets,omitempty" gorm:"foreignKey:SellerID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

type Pet struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"not null;size:255"`
	Species     string    `json:"species" gorm:"not null;size:100"`
	Breed       string    `json:"breed" gorm:"size:100"`
	Age         int       `json:"age" gorm:"check:age >= 0"`
	Price       float64   `json:"price" gorm:"type:decimal(10,2);check:price >= 0"`
	Description string    `json:"description" gorm:"type:text"`
	Available   bool      `json:"available" gorm:"default:true"`
	SellerID    uint      `json:"seller_id" gorm:"not null;index"`
	Seller      *Seller   `json:"seller,omitempty" gorm:"foreignKey:SellerID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateSellerRequest struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type UpdateSellerRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type CreatePetRequest struct {
	Name        string  `json:"name" binding:"required"`
	Species     string  `json:"species" binding:"required"`
	Breed       string  `json:"breed"`
	Age         int     `json:"age"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Available   bool    `json:"available"`
	SellerID    uint    `json:"seller_id" binding:"required"`
}

type UpdatePetRequest struct {
	Name        string  `json:"name"`
	Species     string  `json:"species"`
	Breed       string  `json:"breed"`
	Age         int     `json:"age"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Available   *bool   `json:"available"`
	SellerID    uint    `json:"seller_id"`
}

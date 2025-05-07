package model

import "gorm.io/gorm"

// gorm.Model definition
type Product struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"not null"`
	ProductType string `json:"product_type" gorm:"not null"`
}

func (p *Product) TableName() string {
	return "product"
}

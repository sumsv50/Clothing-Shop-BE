package service

import (
	. "clothing-shop/model"
	"fmt"

	"gorm.io/gorm"
)

type ProductService struct {
	DB *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{DB: db}
}

func (s *ProductService) CreateProduct(product Product) (*Product, error) {
	result := s.DB.Create(&product)
	if result.Error != nil {
		return nil, fmt.Errorf("create product failed: %v", result.Error)
	}

	return &product, nil
}

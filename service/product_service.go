package service

import (
	"clothing-shop/model"
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

func (s *ProductService) GetProducts() ([]*Product, error) {
    var products []*Product
    err := s.DB.Where("is_deleted = ?", false).Find(&products).Error
    if err != nil {
        return nil, fmt.Errorf("create product failed: %v",err)
    }
    return products, nil
}
func (s *ProductService) DeleteProductSoft(id string) error {
    result := s.DB.Model(&model.Product{}).Where("id = ? AND is_deleted = ? ", id, false).Update("is_deleted", true)
    if result.RowsAffected <= 0 {
        return  fmt.Errorf("can not find Product id")
    }
    if result.Error != nil {
        return  result.Error // Product not found, return error
    }
    
    return  nil 
}
func (r *ProductService) Update(product Product) (*Product, error) {
	result := r.DB.Model(&product).Where("is_deleted = ?", false).Updates(&product)
    if result.RowsAffected <= 0 {
        return  nil, fmt.Errorf("can not find Product id")
    }
    if result.Error != nil {
        return  nil, result.Error // Product not found, return error
    }

	return &product, nil
}

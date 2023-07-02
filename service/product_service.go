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

func (s *ProductService) GetProducts() ([]*Product, error) {
    var products []*Product
    err := s.DB.Find(&products).Error
    if err != nil {
        return nil, fmt.Errorf("create product failed: %v",err)
    }
    return products, nil
}
func (s *ProductService) DeleteProductSoft(product Product, id string) ( error) {
    err := s.DB.First(&product, id).Error
    if err != nil {
        return  err // Product not found, return error
    }
    product.IsDelete = true
    err = s.DB.Save(&product).Error
    if err != nil {
        return  err 
    }

    return  nil 
}
func (r *ProductService) Update(product Product) (int64, error) {
	res := r.DB.Save(&product)
	return res.RowsAffected, nil
}

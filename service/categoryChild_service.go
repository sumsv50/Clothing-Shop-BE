package service

import (
	"clothing-shop/model"
	. "clothing-shop/model"
	"fmt"

	"gorm.io/gorm"
)

type CategoryChildService struct {
	DB *gorm.DB
}

func NewCategoryChildService(db *gorm.DB) *CategoryChildService {
	return &CategoryChildService{DB: db}
}

func (s *CategoryChildService) CreateCategoryChild(categoryChild CategoryChilds) (*CategoryChilds, error) {
	result := s.DB.Create(&categoryChild)
	if result.Error != nil {
		return nil, fmt.Errorf("create CategoryChild failed: %v", result.Error)
	}

	return &categoryChild, nil
}

func (s *CategoryChildService) GetCategoryChilds() ([]*CategoryChilds, error) {
    var categoryChilds []*CategoryChilds
    err := s.DB.Where("is_deleted = ?", false).Find(&categoryChilds).Error
    if err != nil {
        return nil, fmt.Errorf("create CategoryChild failed: %v",err)
    }
    return categoryChilds, nil
}
func (s *CategoryChildService) DeleteCategoryChildSoft(id string) error {
    result := s.DB.Model(&model.CategoryChilds{}).Where("id = ? AND is_deleted = ? ", id, false).Update("is_deleted", true)
    if result.RowsAffected <= 0 {
        return  fmt.Errorf("can not find CategoryChild id")
    }
    if result.Error != nil {
        return  result.Error // CategoryChild not found, return error
    }
    
    return  nil 
}
func (r *CategoryChildService) Update(categoryChild CategoryChilds) (*CategoryChilds, error) {
	result := r.DB.Model(&categoryChild).Where("is_deleted = ?", false).Updates(&categoryChild)
    if result.RowsAffected <= 0 {
        return  nil, fmt.Errorf("can not find categoryChild id")
    }
    if result.Error != nil {
        return  nil, result.Error // CategoryChild not found, return error
    }

	return &categoryChild, nil
}

package service

import (
	"clothing-shop/model"
	. "clothing-shop/model"
	"fmt"

	"gorm.io/gorm"
)

type CategoryParentService struct {
	DB *gorm.DB
}

func NewCategoryParentService(db *gorm.DB) *CategoryParentService {
	return &CategoryParentService{DB: db}
}

func (s *CategoryParentService) CreateCategoryParent(categoryParent CategoryParents) (*CategoryParents, error) {
	result := s.DB.Create(&categoryParent)
	if result.Error != nil {
		return nil, fmt.Errorf("create CategoryParent failed: %v", result.Error)
	}

	return &categoryParent, nil
}

func (s *CategoryParentService) GetCategoryParents() ([]*CategoryParents, error) {
	var categoryParents []*CategoryParents
	err := s.DB.Where("is_deleted = ?", false).Order("id").Find(&categoryParents).Error
	if err != nil {
		return nil, fmt.Errorf("create CategoryParent failed: %v", err)
	}
	return categoryParents, nil
}
func (s *CategoryParentService) DeleteCategoryParentSoft(id string) error {
	result := s.DB.Model(&model.CategoryParents{}).Where("id = ? AND is_deleted = ? ", id, false).Update("is_deleted", true)
	if result.RowsAffected <= 0 {
		return fmt.Errorf("can not find CategoryParent id")
	}
	if result.Error != nil {
		return result.Error // CategoryParent not found, return error
	}

	return nil
}
func (r *CategoryParentService) Update(categoryParent CategoryParents) (*CategoryParents, error) {
	result := r.DB.Model(&categoryParent).Where("is_deleted = ?", false).Updates(&categoryParent)
	if result.RowsAffected <= 0 {
		return nil, fmt.Errorf("can not find categoryParent id")
	}
	if result.Error != nil {
		return nil, result.Error // CategoryParent not found, return error
	}

	return &categoryParent, nil
}

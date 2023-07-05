package service

import (
	"clothing-shop/model"
	. "clothing-shop/model"
	"fmt"

	"gorm.io/gorm"
)

type PartnerService struct {
	DB *gorm.DB
}

func NewPartnerService(db *gorm.DB) *PartnerService {
	return &PartnerService{DB: db}
}

func (s *PartnerService) CreatePartner(partner Partner) (*Partner, error) {
	result := s.DB.Create(&partner)
	if result.Error != nil {
		return nil, fmt.Errorf("create Partner failed: %v", result.Error)
	}

	return &partner, nil
}

func (s *PartnerService) GetPartners() ([]*Partner, error) {
    var Partners []*Partner
    err := s.DB.Where("is_deleted = ?", false).Find(&Partners).Error
    if err != nil {
        return nil, fmt.Errorf("create Partner failed: %v",err)
    }
    return Partners, nil
}
func (s *PartnerService) DeletePartnerSoft(id string) error {
    result := s.DB.Model(&model.Partner{}).Where("id = ? AND is_deleted = ? ", id, false).Update("is_deleted", true)
    if result.RowsAffected <= 0 {
        return  fmt.Errorf("can not find Partner id")
    }
    if result.Error != nil {
        return  result.Error // Partner not found, return error
    }
    
    return  nil 
}
func (r *PartnerService) Update(Partner Partner) (*Partner, error) {
	result := r.DB.Model(&Partner).Where("is_deleted = ?", false).Updates(&Partner)
    if result.RowsAffected <= 0 {
        return  nil, fmt.Errorf("can not find Partner id")
    }
    if result.Error != nil {
        return  nil, result.Error // Partner not found, return error
    }

	return &Partner, nil
}

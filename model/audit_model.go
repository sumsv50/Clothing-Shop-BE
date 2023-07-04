package model

import "time"

type Audit struct {
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updated_at"`
	CreatedBy *int       `json:"createdBy" gorm:"column:created_by"`
	UpdatedBy *int       `json:"updatedBy" gorm:"column:updated_by"`
}

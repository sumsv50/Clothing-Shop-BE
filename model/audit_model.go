package model

import "time"

type Audit struct {
	CreatedAt *time.Time `json:"createdAt,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"column:updated_at"`
	CreatedBy *int       `json:"createdBy,omitempty" gorm:"column:created_by"`
	UpdatedBy *int       `json:"updatedBy,omitempty" gorm:"column:updated_by"`
}

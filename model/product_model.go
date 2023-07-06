package model

import (
	pg "github.com/lib/pq"
)

type Product struct {
	Id               *string        `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	ProductCode      *string        `json:"productCode" gorm:"column:productcode"`
	CategoryParentId *int           `json:"categoryParentId" gorm:"column:categoryparentid"`
	CategoryChildId  *int           `json:"categoryChildId" gorm:"column:categorychildid"`
	Title            *string        `json:"title" gorm:"column:title"`
	Details          *string        `json:"details" gorm:"column:details"`
	Images           pg.StringArray `json:"images" gorm:"column:images;type:text[]"`
	OldPrice         *int           `json:"oldPrice" gorm:"column:oldprice"`
	Price            *int           `json:"price" gorm:"column:price"`
	Size             *string        `json:"size" gorm:"column:size"`
	ProductQuality   *string        `json:"productQuality" gorm:"column:productquality"`
	IsDeleted        *bool          `json:"is_deleted" gorm:"column:is_deleted;default:false"`
	Audit            `gorm:"embedded"`
}
type CategoryParents struct {
	Id               *string        `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	Title            *string        `json:"title" gorm:"column:title"`
	IsMenu      	 *string        `json:"ismenu" gorm:"column:ismenu"`
	IsDeleted        *bool           `json:"is_deleted" gorm:"column:is_deleted;default:false"`
}
func (CategoryParents) TableName() string {
	return "categoryparents"
  }
type CategoryChilds struct {
	Id               *string        `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	CategoryParentid *string        `json:"categoryparentid" gorm:"column:categoryparentid"`
	Title            *string        `json:"title" gorm:"column:title"`
	IsMenu      	 *string        `json:"ismenu" gorm:"column:ismenu"`
	IsDeleted        *bool           `json:"is_deleted" gorm:"column:is_deleted;default:false"`
}
func (CategoryChilds) TableName() string {
	return "categorychilds"
  }
type Partner struct {
	Id               *string        `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	Title            *string        `json:"title" gorm:"column:title"`
	Image           *string 		`json:"image" gorm:"column:image"`
	Order            *int        `json:"order" gorm:"column:order"`
	IsDeleted        *bool           `json:"is_deleted" gorm:"column:is_deleted;default:false"`
}

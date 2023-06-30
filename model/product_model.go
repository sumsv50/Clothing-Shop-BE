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
	Image            pg.StringArray `json:"image" gorm:"column:image;type:text[]"`
	OldPrice         *int           `json:"oldPrice" gorm:"column:oldprice"`
	Price            *int           `json:"price" gorm:"column:price"`
	Size             *string        `json:"size" gorm:"column:size"`
	ProductQuality   *string        `json:"productQuality" gorm:"column:productquality"`
}

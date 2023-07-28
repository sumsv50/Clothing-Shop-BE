package model

type CategoryParents struct {
	Id            *string           `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	Title         *string           `json:"title" gorm:"column:title"`
	IsMenu        *string           `json:"ismenu" gorm:"column:ismenu"`
	IsDeleted     *bool             `json:"is_deleted" gorm:"column:is_deleted;default:false"`
	ParentId      *string           `json:"parentId" gorm:"column:parentid"`
	CategoryChild []CategoryParents `json:"child" gorm:"foreignKey:parentid"`
}

func (CategoryParents) TableName() string {
	return "categoryparents"
}

type CategoryChilds struct {
	Id               *string `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	CategoryParentid *string `json:"categoryparentid" gorm:"column:categoryparentid"`
	Title            *string `json:"title" gorm:"column:title"`
	IsMenu           *string `json:"ismenu" gorm:"column:ismenu"`
	IsDeleted        *bool   `json:"is_deleted" gorm:"column:is_deleted;default:false"`
}

func (CategoryChilds) TableName() string {
	return "categorychilds"
}

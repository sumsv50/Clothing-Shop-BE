package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id          *string    `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	Avatar      *string    `json:"avatar" gorm:"column:avatar"`
	Username    *string    `json:"username" gorm:"column:username"`
	Password    *string    `json:"password" gorm:"column:password"`
	Email       *string    `json:"email" gorm:"column:email"`
	PhoneNumber *string    `json:"phoneNumber" gorm:"column:phonenumber"`
	DateOfBirth *time.Time `json:"dob" gorm:"column:dob"`
	Role        *string    `json:"role" gorm:"column:role"`
	IsDeleted   bool       `json:"isDeleted" gorm:"is_deleted"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	hash := string(bytes)
	user.Password = &hash
	return nil
}
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

type LoginReq struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

package service

import (
	. "clothing-shop/model"
	"fmt"

	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) Login(username, password string) (*User, error) {
	user := User{Username: &username}
	result := s.DB.First(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("get user fail: %v", result.Error)
	}
	if result.RowsAffected <= 0 {
		return nil, nil
	}

	err := user.CheckPassword(password)
	if err != nil {
		return nil, nil
	}

	return &user, nil
}

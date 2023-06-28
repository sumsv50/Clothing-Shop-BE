package service

import (
	"database/sql"
	"fmt"
	. "onboarding/exercise1/model"
	"reflect"

	q "github.com/core-go/sql"
)

// type UserService interface {
// 	GetAllUsers() ([]User, error)
// 	GetUserById(id string) (User, error)
// 	CreateUser(user User) (string, error)
// 	UpdateUserByPatch(user map[string]interface{}) (string, error)
// }

type UserService struct {
	DB *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) GetAllUsers() ([]User, error) {
	var users []User

	rows, err := s.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("get all users failed: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Phone, &user.DateOfBirth)
		if err != nil {
			return nil, fmt.Errorf("get all users failed: %v", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("get all users failed: %v", err)
	}

	return users, nil
}

func (s *UserService) GetUserById(id string) (User, error) {
	var user User

	row := s.DB.QueryRow("SELECT * FROM users WHERE id = ?", id)

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Phone, &user.DateOfBirth)

	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, nil
		}
		return user, fmt.Errorf("get user by id failed: %v", err)
	}

	return user, nil
}

func (s *UserService) CreateUser(user User) (User, error) {
	query := "INSERT INTO users (id, username, email, phone, date_of_birth) values (?, ?, ?, ?, ?)"
	_, err := s.DB.Exec(query, user.ID, user.Username, user.Email, user.Phone, user.DateOfBirth)
	if err != nil {
		return User{}, fmt.Errorf("create user failed: %v", err)
	}

	return user, nil
}

func (s *UserService) UpdateUserByPatch(user map[string]interface{}) (User, error) {
	userType := reflect.TypeOf(User{})
	jsonColumnMap := q.MakeJsonColumnMap(userType)
	colMap := q.JSONToColumns(user, jsonColumnMap)
	keys, _ := q.FindPrimaryKeys(userType)
	query, args := q.BuildToPatch("users", colMap, keys, q.BuildParam)
	_, err := s.DB.Exec(query, args...)
	if err != nil {
		return User{}, fmt.Errorf("update user failed: %v", err)
	}
	updatedUser, err := s.GetUserById(fmt.Sprint(user["id"]))

	if err != nil {
		return User{}, fmt.Errorf("cannot get user after updated")
	}

	return updatedUser, nil
}

func (s *UserService) UpdateUserByPut(user User) (User, error) {
	query := "UPDATE users (id, username, email, phone, date_of_birth) values (?, ?, ?, ?, ?)"
	_, err := s.DB.Exec(query, user.ID, user.Username, user.Email, user.Phone, user.DateOfBirth)
	if err != nil {
		return User{}, fmt.Errorf("create user failed: %v", err)
	}

	return user, nil
}

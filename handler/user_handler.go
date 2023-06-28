package handler

import (
	. "clothing-shop/model"
	. "clothing-shop/service"
	"encoding/json"
	"net/http"
	"reflect"

	sv "github.com/core-go/service"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	service UserService
}

func NewUserHandler(s UserService) *UserHandler {
	return &UserHandler{service: s}
}

type RespFormat struct {
	Ok      bool        `json:"ok"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (h *UserHandler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		JSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	JSON(w, http.StatusOK, "", map[string][]User{"users": users})
}

func (h *UserHandler) GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	user, err := h.service.GetUserById(id)
	if err != nil {
		JSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	var resp interface{}
	if !reflect.ValueOf(user).IsZero() {
		resp = user
	}
	JSON(w, http.StatusOK, "", map[string]interface{}{"user": resp})
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if err != nil {
		JSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	createdUser, err := h.service.CreateUser(user)
	if err != nil {
		JSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	JSON(w, http.StatusCreated, "", map[string]interface{}{"user": createdUser})
}

func (h *UserHandler) UpdateUserPatchHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var user User
	userType := reflect.TypeOf(user)
	_, jsonMap, _ := sv.BuildMapField(userType)
	body, err := sv.BuildMapAndStruct(r, &user)
	if err != nil {
		JSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	if len(user.ID) == 0 {
		user.ID = id
	} else if id != user.ID {
		JSON(w, http.StatusBadRequest, "Id not match", nil)
		return
	}
	jsonInfo, err := sv.BodyToJsonMap(r, user, body, []string{"id"}, jsonMap)
	if err != nil {
		JSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	updatedUser, err := h.service.UpdateUserByPatch(jsonInfo)
	if err != nil {
		JSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	var resp interface{}
	if !reflect.ValueOf(updatedUser).IsZero() {
		resp = updatedUser
	}
	JSON(w, http.StatusOK, "", map[string]interface{}{"user": resp})
}

func (h *UserHandler) UpdateUserPutHandler(w http.ResponseWriter, r *http.Request) {
	// userId := mux.Vars(r)["id"]
	// if len(userId) == 0 {
	// 	http.Error(w, "Id cannot be empty", http.StatusBadRequest)
	// 	return
	// }
	// var user User
	// err := json.NewDecoder(r.Body).Decode(&user)
	// defer r.Body.Close()
	// if err != nil {
	// 	JSON(w, http.StatusBadRequest, err.Error(), nil)
	// 	return
	// }
	// user.ID = userId

	// createdUser, err := h.service.UpdateUserByPut(user)
	// if err != nil {
	// 	JSON(w, http.StatusInternalServerError, err.Error(), nil)
	// 	return
	// }
	// JSON(w, http.StatusCreated, "", map[string]interface{}{"user": createdUser})
}

func JSON(w http.ResponseWriter, code int, message string, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	successHttpStatus := []int{
		http.StatusOK,
		http.StatusCreated,
		http.StatusNoContent,
	}

	isOk := false
	for _, value := range successHttpStatus {
		if value == code {
			isOk = true
			break
		}
	}

	return json.NewEncoder(w).Encode(RespFormat{
		Ok:      isOk,
		Message: message,
		Data:    data,
	})
}

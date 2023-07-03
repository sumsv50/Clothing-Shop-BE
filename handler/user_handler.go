package handler

import (
	. "clothing-shop/model"
	"clothing-shop/service"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginReq
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	defer r.Body.Close()
	if err != nil {
		JSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	user, err := h.service.Login(*loginReq.Username, *loginReq.Password)
	if err != nil || user == nil {
		http.Error(w, "", http.StatusUnauthorized)
	}

	token, err := service.GenerateJWT(*user.Id)
	if err != nil {
		JSON(w, http.StatusInternalServerError, "", nil)
		return
	}
	JSON(w, http.StatusOK, "", map[string]interface{}{"token": token})
}

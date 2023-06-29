package handler

import (
	. "clothing-shop/model"
	. "clothing-shop/service"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	service ProductService
}

func NewUserHandler(s ProductService) *UserHandler {
	return &UserHandler{service: s}
}

type RespFormat struct {
	Ok      bool        `json:"ok"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (h *UserHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	defer r.Body.Close()
	if err != nil {
		JSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	createdProduct, err := h.service.CreateProduct(product)
	if err != nil {
		JSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	JSON(w, http.StatusCreated, "", map[string]interface{}{"user": createdProduct})
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

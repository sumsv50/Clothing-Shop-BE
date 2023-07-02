package handler

import (
	. "clothing-shop/model"
	. "clothing-shop/service"
	"encoding/json"
	"github.com/gorilla/mux"
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

func (h *UserHandler) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	var products []*Product
	err := json.NewDecoder(r.Body).Decode(&products)
	defer r.Body.Close()
	if err != nil {
		JSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	products, err = h.service.GetProducts()
	if err != nil {
		JSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	JSON(w, http.StatusCreated, "", map[string]interface{}{"user": products})
}

func (h *UserHandler) SoftDeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	er1 := json.NewDecoder(r.Body).Decode(&product)
	defer r.Body.Close()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	if len(*product.Id) == 0 {
		product.Id = &id
	} else if id != *product.Id {
		http.Error(w, "Id not match", http.StatusBadRequest)
		return
	}
 	er2 := h.service.DeleteProductSoft(product,id)
	if er2 != nil {
		http.Error(w, er2.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusCreated, "", map[string]interface{}{"user": product})
}

func (h *UserHandler) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	er1 := json.NewDecoder(r.Body).Decode(&product)
	defer r.Body.Close()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	if len(*product.Id) == 0 {
		product.Id = &id
	} else if id != *product.Id {
		http.Error(w, "Id not match", http.StatusBadRequest)
		return
	}

	res, er2 := h.service.Update(product)
	if er2 != nil {
		http.Error(w, er2.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusCreated, "", map[string]interface{}{"user": res})
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

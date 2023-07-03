package handler

import (
	. "clothing-shop/model"
	"clothing-shop/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(s service.ProductService) *ProductHandler {
	return &ProductHandler{service: s}
}

type RespFormat struct {
	Ok      bool        `json:"ok"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (h *ProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
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
	JSON(w, http.StatusCreated, "", map[string]interface{}{"product": createdProduct})
}

func (h *ProductHandler) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	var products []*Product

	products, err := h.service.GetProducts()
	if err != nil {
		JSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	JSON(w, http.StatusCreated, "", map[string]interface{}{"products": products})
}

func (h *ProductHandler) SoftDeleteProductHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	er2 := h.service.DeleteProductSoft(id)
	if er2 != nil {
		http.Error(w, er2.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusCreated, "", map[string]interface{}{})
}

func (h *ProductHandler) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
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
	product.Id = &id
	res, er2 := h.service.Update(product)
	if er2 != nil {
		http.Error(w, er2.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusCreated, "", map[string]interface{}{"product": res})
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

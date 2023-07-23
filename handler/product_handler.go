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
	JSON(w, http.StatusOK, "", map[string]interface{}{"products": products})
}

func (h *ProductHandler) SoftDeleteProductHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	er2 := h.service.DeleteProductSoft(id)
	if er2 != nil {
		JSON(w, http.StatusInternalServerError, er2.Error(), nil)
		return
	}
	JSON(w, http.StatusOK, "", map[string]interface{}{})
}

func (h *ProductHandler) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	er1 := json.NewDecoder(r.Body).Decode(&product)
	defer r.Body.Close()
	if er1 != nil {
		JSON(w, http.StatusBadRequest, er1.Error(), nil)
		return
	}
	id := mux.Vars(r)["id"]
	product.Id = &id
	res, er2 := h.service.Update(product)
	if er2 != nil {
		JSON(w, http.StatusInternalServerError, er2.Error(), nil)
		return
	}
	JSON(w, http.StatusOK, "", map[string]interface{}{"product": res})
}

func (h *ProductHandler) GetProductDetailHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	product, er2 := h.service.GetProductDetail(id)
	if er2 != nil {
		JSON(w, http.StatusInternalServerError, er2.Error(), nil)
		return
	}
	JSON(w, http.StatusOK, "", map[string]interface{}{"product": product})
}

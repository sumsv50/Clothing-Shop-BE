package handler

import (
	. "clothing-shop/model"
	. "clothing-shop/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type CategoryChildHandler struct {
	service CategoryChildService
}

func NewCategoryChildHandler(s CategoryChildService) *CategoryChildHandler {
	return &CategoryChildHandler{service: s}
}


func (h *CategoryChildHandler) CreateCategoryChildHandler(w http.ResponseWriter, r *http.Request) {
	var categoryChild CategoryChilds
	err := json.NewDecoder(r.Body).Decode(&categoryChild)
	defer r.Body.Close()
	if err != nil {
		JSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	res, err := h.service.CreateCategoryChild(categoryChild)
	if err != nil {
		JSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	JSON(w, http.StatusCreated, "", map[string]interface{}{"CategoryChilds": res})
}

func (h *CategoryChildHandler) GetCategoryChildsHandler(w http.ResponseWriter, r *http.Request) {
	var CategoryChilds []*CategoryChilds

	CategoryChilds, err := h.service.GetCategoryChilds()
	if err != nil {
		JSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	JSON(w, http.StatusCreated, "", map[string]interface{}{"CategoryChilds": CategoryChilds})
}

func (h *CategoryChildHandler) SoftDeleteCategoryChildHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	er2 := h.service.DeleteCategoryChildSoft(id)
	if er2 != nil {
		http.Error(w, er2.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusCreated, "", map[string]interface{}{})
}

func (h *CategoryChildHandler) UpdateCategoryChildHandler(w http.ResponseWriter, r *http.Request) {
	var categoryChild CategoryChilds
	er1 := json.NewDecoder(r.Body).Decode(&categoryChild)
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
	categoryChild.Id = &id
	res, er2 := h.service.Update(categoryChild)
	if er2 != nil {
		http.Error(w, er2.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusCreated, "", map[string]interface{}{"categoryChild": res})
}


package handler

import (
	. "clothing-shop/model"
	. "clothing-shop/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type CategoryParentHandler struct {
	service CategoryParentService
}

func NewCategoryParentHandler(s CategoryParentService) *CategoryParentHandler {
	return &CategoryParentHandler{service: s}
}


func (h *CategoryParentHandler) CreateCategoryParentHandler(w http.ResponseWriter, r *http.Request) {
	var categoryParent CategoryParents
	err := json.NewDecoder(r.Body).Decode(&categoryParent)
	defer r.Body.Close()
	if err != nil {
		JSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	res, err := h.service.CreateCategoryParent(categoryParent)
	if err != nil {
		JSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	JSON(w, http.StatusCreated, "", map[string]interface{}{"categoryParent": res})
}

func (h *CategoryParentHandler) GetCategoryParentsHandler(w http.ResponseWriter, r *http.Request) {
	var categoryParents []*CategoryParents

	categoryParents, err := h.service.GetCategoryParents()
	if err != nil {
		JSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	JSON(w, http.StatusCreated, "", map[string]interface{}{"categoryParents": categoryParents})
}

func (h *CategoryParentHandler) SoftDeleteCategoryParentHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	er2 := h.service.DeleteCategoryParentSoft(id)
	if er2 != nil {
		http.Error(w, er2.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusCreated, "", map[string]interface{}{})
}

func (h *CategoryParentHandler) UpdateCategoryParentHandler(w http.ResponseWriter, r *http.Request) {
	var categoryParent CategoryParents
	er1 := json.NewDecoder(r.Body).Decode(&categoryParent)
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
	categoryParent.Id = &id
	res, er2 := h.service.Update(categoryParent)
	if er2 != nil {
		http.Error(w, er2.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusCreated, "", map[string]interface{}{"categoryParent": res})
}


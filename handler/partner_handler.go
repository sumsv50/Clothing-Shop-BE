package handler

import (
	. "clothing-shop/model"
	. "clothing-shop/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type PartnerHandler struct {
	service PartnerService
}

func NewPartnerHandler(s PartnerService) *PartnerHandler {
	return &PartnerHandler{service: s}
}


func (h *PartnerHandler) CreatePartnerHandler(w http.ResponseWriter, r *http.Request) {
	var partner Partner
	err := json.NewDecoder(r.Body).Decode(&partner)
	defer r.Body.Close()
	if err != nil {
		JSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	res, err := h.service.CreatePartner(partner)
	if err != nil {
		JSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	JSON(w, http.StatusCreated, "", map[string]interface{}{"partner": res})
}

func (h *PartnerHandler) GetPartnersHandler(w http.ResponseWriter, r *http.Request) {
	var partners []*Partner

	partners, err := h.service.GetPartners()
	if err != nil {
		JSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	JSON(w, http.StatusCreated, "", map[string]interface{}{"partners": partners})
}

func (h *PartnerHandler) SoftDeletePartnerHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	er2 := h.service.DeletePartnerSoft(id)
	if er2 != nil {
		http.Error(w, er2.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusCreated, "", map[string]interface{}{})
}

func (h *PartnerHandler) UpdatePartnerHandler(w http.ResponseWriter, r *http.Request) {
	var product Partner
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


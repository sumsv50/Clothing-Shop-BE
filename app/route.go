package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Route() {
	app := NewApp()
	const apiPath = "/v1/clothing/api"
	const product = "/products"
	const auth = "/auth"
	const partner = "/partners"
	myRouter := mux.NewRouter().StrictSlash(true)
	protectedRouter := myRouter.NewRoute().Subrouter()
	// Turn off authentication admin APIs
	// protectedRouter.Use(middleware.AuthenticationMiddleware)

	// Protected product APIs
	protectedRouter.HandleFunc(apiPath+product, app.ProductHandler.CreateProductHandler).Methods("POST")
	protectedRouter.HandleFunc(apiPath+product+"/{id}", app.ProductHandler.SoftDeleteProductHandler).Methods("DELETE")
	protectedRouter.HandleFunc(apiPath+product+"/{id}", app.ProductHandler.UpdateProductHandler).Methods("PATCH")

	// Product APIs
	myRouter.HandleFunc(apiPath+product, app.ProductHandler.GetProductsHandler).Methods("GET")

	// Auth APIs
	myRouter.HandleFunc(apiPath+auth+"/local", app.UserHandler.Login).Methods("POST")

	// Partner API
	myRouter.HandleFunc(apiPath+partner, app.PartnerHandler.CreatePartnerHandler).Methods("POST")
	myRouter.HandleFunc(apiPath+partner, app.PartnerHandler.GetPartnersHandler).Methods("GET")
	myRouter.HandleFunc(apiPath+partner+"/{id}", app.PartnerHandler.SoftDeletePartnerHandler).Methods("DELETE")
	myRouter.HandleFunc(apiPath+partner+"/{id}", app.PartnerHandler.UpdatePartnerHandler).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

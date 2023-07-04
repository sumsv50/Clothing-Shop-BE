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

	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

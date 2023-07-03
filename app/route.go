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
	// Products API
	myRouter.HandleFunc(apiPath+product, app.ProductHandler.CreateProductHandler).Methods("POST")
	myRouter.HandleFunc(apiPath+product, app.ProductHandler.GetProductsHandler).Methods("GET")
	myRouter.HandleFunc(apiPath+product+"/{id}", app.ProductHandler.SoftDeleteProductHandler).Methods("DELETE")
	myRouter.HandleFunc(apiPath+product+"/{id}", app.ProductHandler.UpdateProductHandler).Methods("PATCH")

	// Auth API
	myRouter.HandleFunc(apiPath+auth+"/local", app.UserHandler.Login).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

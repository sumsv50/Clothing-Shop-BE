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
	myRouter := mux.NewRouter().StrictSlash(true)
	// Products API
	myRouter.HandleFunc( apiPath + product , app.UserHandler.CreateProductHandler).Methods("POST")
	myRouter.HandleFunc( apiPath + product , app.UserHandler.GetProductsHandler).Methods("GET")
	myRouter.HandleFunc( apiPath + product + "/detele/{id}", app.UserHandler.SoftDeleteProductHandler).Methods("PUT")
	myRouter.HandleFunc( apiPath + product + "/update/{id}", app.UserHandler.UpdateProductHandler).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

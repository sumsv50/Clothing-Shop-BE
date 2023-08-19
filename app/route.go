package app

import (
	"clothing-shop/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Route() {
	app := NewApp()
	const apiPath = "/v1/clothing/api"
	const product = "/products"
	const auth = "/auth"
	const partner = "/partners"
	const category = "/category"
	const parent = "/parent"
	const child = "/child"

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodDelete},
	})

	myRouter := mux.NewRouter().StrictSlash(true)
	protectedRouter := myRouter.NewRoute().Subrouter()
	// Turn off authentication admin APIs
	protectedRouter.Use(middleware.AuthenticationMiddleware)

	// Protected product APIs
	protectedRouter.HandleFunc(apiPath+product, app.ProductHandler.CreateProductHandler).Methods("POST")
	protectedRouter.HandleFunc(apiPath+product+"/{id}", app.ProductHandler.SoftDeleteProductHandler).Methods("DELETE")
	protectedRouter.HandleFunc(apiPath+product+"/{id}", app.ProductHandler.UpdateProductHandler).Methods("PATCH")

	// Product APIs
	myRouter.HandleFunc(apiPath+product, app.ProductHandler.GetProductsHandler).Methods("GET")
	protectedRouter.HandleFunc(apiPath+product+"/{id}", app.ProductHandler.GetProductDetailHandler).Methods("GET")

	// Auth APIs
	myRouter.HandleFunc(apiPath+auth+"/local", app.UserHandler.Login).Methods("POST")

	// Partner API
	protectedRouter.HandleFunc(apiPath+partner, app.PartnerHandler.CreatePartnerHandler).Methods("POST")
	myRouter.HandleFunc(apiPath+partner, app.PartnerHandler.GetPartnersHandler).Methods("GET")
	protectedRouter.HandleFunc(apiPath+partner+"/{id}", app.PartnerHandler.SoftDeletePartnerHandler).Methods("DELETE")
	protectedRouter.HandleFunc(apiPath+partner+"/{id}", app.PartnerHandler.UpdatePartnerHandler).Methods("PATCH")

	// category parent API
	protectedRouter.HandleFunc(apiPath+category+parent, app.CategoryParentHandler.CreateCategoryParentHandler).Methods("POST")
	myRouter.HandleFunc(apiPath+category+parent, app.CategoryParentHandler.GetCategoryParentsHandler).Methods("GET")
	protectedRouter.HandleFunc(apiPath+category+parent+"/{id}", app.CategoryParentHandler.SoftDeleteCategoryParentHandler).Methods("DELETE")
	protectedRouter.HandleFunc(apiPath+category+parent+"/{id}", app.CategoryParentHandler.UpdateCategoryParentHandler).Methods("PATCH")

	// category child API
	myRouter.HandleFunc(apiPath+category+child, app.CategoryChildHandler.CreateCategoryChildHandler).Methods("POST")
	myRouter.HandleFunc(apiPath+category+child, app.CategoryChildHandler.GetCategoryChildsHandler).Methods("GET")
	myRouter.HandleFunc(apiPath+category+child+"/{id}", app.CategoryChildHandler.SoftDeleteCategoryChildHandler).Methods("DELETE")
	myRouter.HandleFunc(apiPath+category+child+"/{id}", app.CategoryChildHandler.UpdateCategoryChildHandler).Methods("PATCH")

	handler := cors.Handler(myRouter)
	log.Fatal(http.ListenAndServe(":3000", handler))
}

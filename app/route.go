package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Route() {
	app := NewApp()

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/products", app.UserHandler.CreateProductHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

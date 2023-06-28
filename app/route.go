package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Route() {
	app := NewApp()

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", app.UserHandler.GetAllUsersHandler).Methods("GET")
	myRouter.HandleFunc("/users/{id}", app.UserHandler.GetUserByIdHandler).Methods("GET")
	myRouter.HandleFunc("/users", app.UserHandler.CreateUserHandler).Methods("POST")
	myRouter.HandleFunc("/users/{id}", app.UserHandler.UpdateUserPatchHandler).Methods("PATCH")
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

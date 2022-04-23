package routes

import (
	"github.com/api_app/handlers"
	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/users/update/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/delete/{id}", handlers.DeleteUser).Methods("DELETE")
}

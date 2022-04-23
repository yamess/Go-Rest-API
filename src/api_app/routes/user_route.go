package routes

import (
	"github.com/api_app/services"
	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	r.HandleFunc("/users", services.GetUsers).Methods("GET")
	r.HandleFunc("/users", services.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", services.GetUser).Methods("GET")
	r.HandleFunc("/users/update/{id}", services.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/delete/{id}", services.DeleteUser).Methods("DELETE")
}

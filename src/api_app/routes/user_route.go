package routes

import (
	"github.com/api_app/handlers"
	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	userRoute := r.PathPrefix("/users").Subrouter()

	userRoute.HandleFunc("/all", handlers.AllUsers).Methods("GET")
	userRoute.HandleFunc("/create", handlers.NewUser).Methods("POST")
	userRoute.HandleFunc("/{id}", handlers.SingleUser).Methods("GET")
	userRoute.HandleFunc("/update/{id}", handlers.UpdateUser).Methods("PUT")
	userRoute.HandleFunc("/delete/{id}", handlers.DeleteUser).Methods("DELETE")
}

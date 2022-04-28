package routes

import (
	middleware "github.com/api_app/middlewares"
	"github.com/api_app/services"
	"github.com/gorilla/mux"
	"net/http"
)

func UserRoutes(r *mux.Router) {

	getAllRouter := r.Methods(http.MethodGet).Subrouter()
	getAllRouter.HandleFunc("/users", services.GetUsers)
	//getRouter.Use(middleware.SetupCORS, middleware.Logging)

	postRouter := r.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/users", services.CreateUser)
	postRouter.Use(middleware.Validator)

	getRouter := r.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/users/{id}", services.GetUserByID)
	//r.HandleFunc("/users/{id}", services.GetUserByID).Methods("GET")

	updateRouter := r.Methods(http.MethodPatch).Subrouter()
	updateRouter.HandleFunc("/users/update/{id}", services.UpdateUser)
	//r.HandleFunc("/users/update/{id}", middleware.ChainMiddleware(services.UpdateUser)).Methods("PUT")

	deleteRouter := r.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/users/delete/{id}", services.DeleteUser)
	//r.HandleFunc("/users/delete/{id}", middleware.ChainMiddleware(services.DeleteUser)).Methods("DELETE")
}

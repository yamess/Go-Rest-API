package routes

import (
	"github.com/api_app/services"
	"github.com/gorilla/mux"
	"net/http"
)

func AuthRoutes(r *mux.Router) {
	signInRouter := r.Methods(http.MethodPost).Subrouter()
	signInRouter.HandleFunc("/signing", services.SignIn)
}

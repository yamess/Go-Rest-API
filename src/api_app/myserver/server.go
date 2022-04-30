package myserver

import (
	"github.com/api_app/constants"
	middleware "github.com/api_app/middlewares"
	"github.com/api_app/routes"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func Server() {
	baseRouter := mux.NewRouter()
	baseRouter.Use(middleware.SetupCORS, middleware.Logging)

	allRouter := baseRouter.PathPrefix("/api").Subrouter()

	routes.UserRoutes(allRouter)
	routes.AuthRoutes(allRouter)
	allRouter.Use(middleware.SetupCORS, middleware.Logging)

	// CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	srv := &http.Server{
		Handler:      ch(allRouter),
		Addr:         constants.Host,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

package myserver

import (
	"github.com/api_app/constants"
	middleware "github.com/api_app/middlewares"
	"github.com/api_app/routes"
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
	allRouter.Use(middleware.SetupCORS, middleware.Logging)

	srv := &http.Server{
		Handler:      allRouter,
		Addr:         constants.Host,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

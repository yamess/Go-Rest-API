package myserver

import (
	"github.com/api_app/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func Server() {
	baseRouter := mux.NewRouter()

	allRouter := baseRouter.PathPrefix("/api").Subrouter()

	routes.UserRoutes(allRouter)

	srv := &http.Server{
		Handler:      allRouter,
		Addr:         ":8081",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

package middleware

import (
	"context"
	"fmt"
	"github.com/api_app/models"
	"log"
	"net/http"
)

type KeyUser struct{}

func Validator(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := user.FromJSON(r.Body)
		if err != nil {
			log.Println("Error while decoding data")
		}
		err = user.Validate()
		if err != nil {
			log.Printf("Invalid data\n")
			log.Println(err.Error())
			errorString := fmt.Sprintf("Invalid data \n%s", err.Error())
			http.Error(w, errorString, http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), KeyUser{}, user)
		r = r.WithContext(ctx)
		f.ServeHTTP(w, r)
	})
}

func Logging(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		f.ServeHTTP(w, r)
	})
}

func SetupCORS(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length,"+
			"Accept-Encoding, X-CSRF-Token, Authorization",
		)
		f.ServeHTTP(w, r)
	})
}

package middlewares

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do some stuff here

		log.Println(r.RequestURI)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

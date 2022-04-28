package middleware

import "net/http"

type Middleware func(handlerFunc http.HandlerFunc) http.HandlerFunc

func ChainMiddleware(h http.HandlerFunc, m ...Middleware) http.HandlerFunc {
	if len(m) < 1 {
		return h
	}
	wrapped := h

	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}
	return wrapped
}

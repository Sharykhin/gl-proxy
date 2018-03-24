package middleware

import "net/http"

type (
	// Middleware defines new type for http.Handler wrapper
	Middleware func(handler http.Handler) http.Handler
)

//Chain is a helper function that allows to chain multiple middlewares
func Chan(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

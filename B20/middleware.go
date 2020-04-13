package main

import "net/http"

//USERNAME is
const USERNAME = "batman"

//PASSWORD is
const PASSWORD = "secret"

type CustomMux struct {
	http.ServeMux
	middleware []func(next http.Handler) http.Handler
}

func (c *CustomMux) RegisterMiddleware(next func(next http.Handler) http.Handler)  {
	c.middleware = append(c.middleware, next)
}

func (c *CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	var current http.Handler = &c.ServeMux

	for _, next := range c.middleware {
		current = next(current)
	}

	current.ServeHTTP(w, r)
}

//MiddlewareAuth is
func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte(`something went wrong`))
			return
		}

		isValid := (username == USERNAME) && (password == PASSWORD)
		if !isValid {
			w.Write([]byte(`wrong username/password`))
			return
		}

		next.ServeHTTP(w, r)
	})
}

//MiddlewareAllowOnlyGet is
func MiddlewareAllowOnlyGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.Write([]byte("Only GET is allowed"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

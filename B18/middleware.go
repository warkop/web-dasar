package main

import "net/http"

//USERNAME is
const USERNAME = "batman"

//PASSWORD is
const PASSWORD = "secret"

//Auth is
func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte(`something went wrong`))
		return false
	}

	isValid := (username == USERNAME) && (password == PASSWORD)
	if !isValid {
		w.Write([]byte(`wrong username/password`))
		return false
	}

	return true
}

//AllowOnlyGET is
func AllowOnlyGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte(`Only GET is allowed`))
		return false
	}

	return true
}

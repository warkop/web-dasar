package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func routeIndex(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	var tmpl = template.Must(template.ParseFiles("view.html"))
	var err = tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func routeSubmitPost(w http.ResponseWriter, r *http.Request)  {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main()  {
	http.HandleFunc("/", routeIndex)
	http.HandleFunc("/process", routeSubmitPost)

	fmt.Println("server running at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

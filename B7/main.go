package main

import (
	"fmt"
	"net/http"
	"text/template"
)

//Superhero is
type Superhero struct {
	Name    string
	Alias   string
	Friends []string
}

//SayHello is
func (s Superhero) SayHello(from string, message string) string {
	return fmt.Sprintf("%s said: \"%s\"", from, message)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var Person = Superhero{
			Name:    "Bruce Wayne",
			Alias:   "Batman",
			Friends: []string{"Superman", "Flash", "Green Lantern"},
		}

		var tmpl = template.Must(template.ParseFiles("view.html"))
		if err := tmpl.Execute(w, Person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server running at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

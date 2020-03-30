package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//Info is a...
type Info struct {
	Affiliation string
	Address     string
}

//Person is a...
type Person struct {
	Name    string
	Gender  string
	Hobbies []string
	Info    Info
}

//GetAffiliationDetailInfo is....
func (t Info) GetAffiliationDetailInfo() string {
	return "have 31 divisions"
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var person = Person{
			Name:    "Bruce Wayne",
			Gender:  "male",
			Hobbies: []string{"Reading books", "Soccer", "Fishing"},
			Info:    Info{"Wayne Enterprises", "Gotham City"},
		}

		var tmpl = template.Must(template.ParseFiles("views/view.html"))
		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server running at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

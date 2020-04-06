package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/novalagung/gubrak"
)

//M is
type M map[string]interface{}

var cookieName = "CookieData"

//ActionIndex is
func ActionIndex(w http.ResponseWriter, r *http.Request) {
	cookieName := "CookieData"

	c := &http.Cookie{}

	if storeCookie, _ := r.Cookie(cookieName); storeCookie != nil {
		c = storeCookie
	}

	if c.Value == "" {
		c = &http.Cookie{}
		c.Name = cookieName
		c.Value = gubrak.RandomString(32)
		c.Expires = time.Now().Add(5 * time.Minute)
		http.SetCookie(w, c)
	}

	w.Write([]byte(c.Value))
}

//ActionDelete is
func ActionDelete(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{}
	c.Name = cookieName
	c.Expires = time.Unix(0, 0)
	c.MaxAge = -1
	http.SetCookie(w, c)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func main() {
	http.HandleFunc("/", ActionIndex)
	http.HandleFunc("/delete", ActionDelete)

	fmt.Println("server running at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

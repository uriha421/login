package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

// a template with its filename
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})

	err := t.templ.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	server := http.Server{
		Addr: "localhost:8080",
	}

	http.Handle("/signup", &templateHandler{filename: "signup.html"})
	http.HandleFunc("/setuser", setUser)

	http.Handle("/signin", &templateHandler{filename: "signin.html"})
	http.HandleFunc("/authuser", authUser)

	server.ListenAndServe()
}

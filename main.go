package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/login.html"))
	t.Execute(w, nil)
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := getUserByEmail(r.PostFormValue("email"))

	if user.Password != r.PostFormValue("password") {
		fmt.Println("invalid password")
	}

	session, err := user.createSession()
	if err != nil {
		log.Fatal(err)
	}

	cookie := http.Cookie{
		Name:     "_cookie",
		Value:    session.Uuid,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/", 302)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	session, err := session(w, r)
	if err != nil {
		fmt.Println("error")
		http.Redirect(w, r, "/login", 302)
	}

	fmt.Println(session)

	t := template.Must(template.ParseFiles("templates/index.html"))
	t.Execute(w, nil)
}

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/auth", authHandler)

	server.ListenAndServe()
}

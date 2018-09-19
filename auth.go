package main

import (
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/login.html")
	t.Execute(w, nil)
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	// get a map of email and password
	err = r.ParseForm()
	if err != nil {
		// fail
		http.Redirect(w, r, "/login", 302)
	}

	// get a user from database by email
	user, err := getUserByEmail(r.PostFormValue("email"))
	if err != nil {
		// fail
		http.Redirect(w, r, "/login", 302)
	}

	// compare two hashes of the passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.PostFormValue("password")))
	if err != nil {
		// fail
		http.Redirect(w, r, "/login", 302)
	}

	// create a session
	session, err := user.createSession()
	if err != nil {
		// fail
		http.Redirect(w, r, "/login", 302)
	}

	// create and set a cookie
	cookie := http.Cookie{
		Name:     "_cookie",
		Value:    session.Uuid,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	// succeed in authentication
	http.Redirect(w, r, "/", 302)
}

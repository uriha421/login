package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func authUser(w http.ResponseWriter, r *http.Request) {
	// parse the request body
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	// get a User from database
	user, err := getUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Fatal(err)
	}

	// compare two hashed passwords
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(r.PostFormValue("password")))
	if err != nil {
		log.Fatal(err)
	}

	// create a session
	session, err := user.createSession()
	if err != nil {
		log.Fatal(err)
	}

	// set a cookie
	cookie := http.Cookie{
		Name:     "login_cookie",
		Value:    string(session.Uuid),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	fmt.Fprintln(w, "you have succeeded in sign in")
}

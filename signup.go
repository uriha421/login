package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

func setUser(w http.ResponseWriter, r *http.Request) {
	// parse the request body
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	// generate hash(password)
	password := []byte(r.PostFormValue("password"))
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	// create a user
	u := User{
		Email:     r.PostFormValue("email"),
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}
	u.createUser()

	fmt.Fprintln(w, "you have succeeded in sign up")
}

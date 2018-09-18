package main

import (
  "time"
)

func signupHandler(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("templates/signup.html")
  t.Execute(w, nil)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
  var err error

  // get a map of name, email, and password
  err = r.ParseForm()
  if err != nil {
    // fail
    http.Redirect(w, r, "/login", 302)
  }

  // check email
  b, err = notYetRegistered(r.PostFormValue("email"))
  if !b || err != nil {
    // fail
    http.Redirect(w, r, "/login", 302)
  }

  // get a hashed password
  hashPassword, err := bcrypt.GenerateFromPassword([]byte(r.PostFormValue("password")), bcrypt.DefaultCost)
  if err != nil {
    log.Fatal(err)
  }

  // create a User structure
  user := User{
    Uuid: createUUID(),
    Name: r.PostFormValue("name"),
    Email: r.PostFormValue("email"),
    Password: hashPassword,
    CreatedAt: time.Now(),
  }

  // create a user in database
  user.Create()

  // succeed in signup
  http.Redirect(w, r, "/", 302)
}

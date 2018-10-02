package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

func session(w http.ResponseWriter, r *http.Request) (s Session, err error) {
	// get the client's cookie
	cookie, err := r.Cookie("login_cookie")
	if err != nil {
		return
	}

	// make sure whether session exists or not
	s = Session{Uuid: []byte(cookie.Value)}
	ok, err := s.check()
	if err != nil {
		return
	}
	if !ok {
		err = errors.New("Invalid Session")
	}

	return
}

func showTodos(w http.ResponseWriter, r *http.Request) {
	// check the session
	s, err := session(w, r)
	if err != nil {
		log.Fatal(err)
	}

	// get a todos' list
	Todos, err := getTodos(s.UserId)

	// make a template
	t := template.Must(template.ParseFiles("templates/todos.html"))
	t.Execute(w, Todos)
}

func addTodos(w http.ResponseWriter, r *http.Request) {
	// check the session
	s, err := session(w, r)
	if err != nil {
		log.Fatal(err)
	}

	// parse the request body
	err = r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	// parse completed
	completed, err := strconv.Atoi(r.PostFormValue("completed"))
	if err != nil {
		log.Fatal(err)
	}

	// parse due
	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(r.PostFormValue("due"), layout)
	if err != nil {
		log.Fatal(err)
	}

	// create a todo
	t := Todo{
		UserId:    s.UserId,
		Body:      r.PostFormValue("body"),
		Completed: completed,
		Due:       parsedTime,
		CreatedAt: time.Now(),
	}
	t.createTodo()

	http.Redirect(w, r, "/showtodos", 302)
}

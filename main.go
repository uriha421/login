package main

import (
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/auth", authHandler)
  http.HandleFunc("/signup", signupHandler)
  http.HandleFunc("/register", registerHandler)
	server.ListenAndServe()
}

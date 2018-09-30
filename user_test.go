package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	u := User{
		Email:     "example@gmail.com",
		Password:  []byte("password"),
		CreatedAt: time.Now(),
	}

	err := u.createUser()
	if err != nil {
		t.Fatalf("failed to createUser %#v", err)
	}
}

func TestGetUserByEmail(t *testing.T) {
	u := User{
		Email:     "example@gmail.com",
		Password:  []byte("password"),
		CreatedAt: time.Now(),
	}

	err := u.createUser()
	if err != nil {
		t.Fatalf("failed to createUser %#v", err)
	}

	user, err := getUserByEmail("example@gmail.com")
	if err != nil {
		t.Fatalf("failed to getUserByEmail %#v", err)
	}

	fmt.Println(user)
}

func TestCreateSession(t *testing.T) {
	u := User{
		Email:     "example@gmail.com",
		Password:  []byte("password"),
		CreatedAt: time.Now(),
	}

	err := u.createUser()
	if err != nil {
		t.Fatalf("failed to createUser %#v", err)
	}

	session, err := u.createSession()
	if err != nil {
		t.Fatalf("failed to createSession %#v", err)
	}

	fmt.Println(session)
}

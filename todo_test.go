package main

import (
	"testing"
	"time"
)

func TestCreateTodo(t *testing.T) {
	u := User{
		Email:     "example@gmail.com",
		Password:  []byte("password"),
		CreatedAt: time.Now(),
	}

	err := u.createUser()
	if err != nil {
		t.Fatalf("failed to createUser %#v", err)
	}

	todo := Todo{
	    UserId: u.Id,
	    Content: "I MUST SAY HELLO WORLD",
	    Completed: 0,
	    Due: time.Now(),
	    CreatedAt: time.Now(),
	}

	err = todo.createTodo()
	if err != nil {
		t.Fatalf("failed to createTodo %#v", err)
	}
}


func TestPutContent(t *testing.T) {
	u := User{
		Email:     "example@gmail.com",
		Password:  []byte("password"),
		CreatedAt: time.Now(),
	}

	err := u.createUser()
	if err != nil {
		t.Fatalf("failed to createUser %#v", err)
	}

	todo := Todo{
	    UserId: u.Id,
	    Content: "I MUST SAY HELLO WORLD",
	    Completed: 0,
	    Due: time.Now(),
	    CreatedAt: time.Now(),
	}

	err = todo.createTodo()
	if err != nil {
		t.Fatalf("failed to createTodo %#v", err)
	}

	todo.putContent("I HAVE TO SAY HELLO WORLD")
}

func TestPutCompleted(t *testing.T) {
	u := User{
		Email:     "example@gmail.com",
		Password:  []byte("password"),
		CreatedAt: time.Now(),
	}

	err := u.createUser()
	if err != nil {
		t.Fatalf("failed to createUser %#v", err)
	}

	todo := Todo{
	    UserId: u.Id,
	    Content: "I MUST SAY HELLO WORLD",
	    Completed: 0,
	    Due: time.Now(),
	    CreatedAt: time.Now(),
	}

	err = todo.createTodo()
	if err != nil {
		t.Fatalf("failed to createTodo %#v", err)
	}

	todo.putCompleted(50)
}


func TestPutDue(t *testing.T) {
	u := User{
		Email:     "example@gmail.com",
		Password:  []byte("password"),
		CreatedAt: time.Now(),
	}

	err := u.createUser()
	if err != nil {
		t.Fatalf("failed to createUser %#v", err)
	}

	todo := Todo{
	    UserId: u.Id,
	    Content: "I MUST SAY HELLO WORLD",
	    Completed: 0,
	    Due: time.Now(),
	    CreatedAt: time.Now(),
	}

	err = todo.createTodo()
	if err != nil {
		t.Fatalf("failed to createTodo %#v", err)
	}

	time.Sleep(100 * time.Millisecond)

	todo.putDue(time.Now())
}



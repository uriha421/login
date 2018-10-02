package main

import (
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type User struct {
	Id        int
	Email     string
	Password  []byte
	CreatedAt time.Time
}

type Session struct {
	Id        int
	Uuid      []byte
	UserId    int
	CreatedAt time.Time
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

func (u *User) createUser() (err error) {
	err = Db.QueryRow("INSERT into users (email, password, created_at) values ($1, $2, $3) returning id", u.Email, u.Password, u.CreatedAt).Scan(&u.Id)
	return
}

func getUserByEmail(email string) (u User, err error) {
	err = Db.QueryRow("SELECT * from users WHERE email = $1", email).Scan(&u.Id, &u.Email, &u.Password, &u.CreatedAt)
	return
}

func (u *User) createSession() (s Session, err error) {
	err = Db.QueryRow("INSERT into sessions (uuid, user_id, created_at) values ($1, $2, $3) returning id, uuid, user_id, created_at", uuid.New(), u.Id, time.Now()).Scan(&s.Id, &s.Uuid, &s.UserId, &s.CreatedAt)
	return
}

func (s *Session) check() (b bool, err error) {
	err = Db.QueryRow("SELECT id, uuid, user_id, created_at FROM sessions WHERE uuid = $1", s.Uuid).Scan(&s.Id, &s.Uuid, &s.UserId, &s.CreatedAt)

	if err != nil {
	    b = false
	    return
	}

	if s.Id != 0 {
	    b = true
	    return
	}

	return
}

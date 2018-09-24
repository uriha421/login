package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	Admin     int
	CreatedAt time.Time
}

type Session struct {
	Id        int
	Uuid      string
	UserId    int
	Admin     int
	CreatedAt time.Time
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func getUserByEmail(email string) (u User, err error) {
	err = Db.QueryRow("SELECT id, uuid, name, email, password, admin, created_at FROM users WHERE email = $1", email).Scan(&u.Id, &u.Uuid, &u.Name, &u.Email, &u.Password, &u.Admin, &u.CreatedAt)
	return
}

func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

func (u *User) createSession() (s Session, err error) {
	err = Db.QueryRow("INSERT into sessions (uuid, user_id, admin, created_at) values ($1, $2, $3, $4) returning id, uuid, user_id, admin, created_at", createUUID(), u.Id, u.Admin, time.Now()).Scan(&s.Id, &s.Uuid, &s.UserId, &s.Admin, &s.CreatedAt)
	return
}

func session(w http.ResponseWriter, r *http.Request) (s Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		return
	}

	s = Session{Uuid: cookie.Value}
	ok, _ := s.Check()
	if !ok {
		err = errors.New("invalid session")
	}

	return
}

func (s *Session) Check() (b bool, err error) {
	err = Db.QueryRow("SELECT id, uuid, user_id, admin, created_at FROM sessions WHERE uuid = $1", s.Uuid).
		Scan(&s.Id, &s.Uuid, &s.UserId, &s.Admin, &s.CreatedAt)
	if err != nil {
		b = false
		return
	}
	if s.Id != 0 {
		b = true
	}
	return
}

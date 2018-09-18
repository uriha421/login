package main

import (
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

func getUserByEmail(email string) (u User, err error) {
	u := User{}
	err = dB.QueryRow("SELECT id, uuid, name, password, created_at FROM users WHERE email = $1 returning id, uuid, name, password, created_at", email).Scan(&u.Id, &u.Uuid, &u.Name, &u.Password, &u.CreatedAt)
	return
}

func notYetRegistered(email string) (b bool, err error) {
	// TODO
	// err = dB.QueryRow("select count(*) from users where email = $1", email)
	b = true
	err = nil
	return
}

func (u *User) createUser() (err error) {
	err = dB.QueryRow("INSERT into users (uuid, name, email, password, created_at) values ($1, $2, $3, $4, $5) returning id",
		u.Uuid, u.Name, u.Email, u.Password, u.CreatedAt).Scan(&u.Id)
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
	err = dB.QueryRow("INSERT into sessions (uuid, email, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, email, user_id, created_at", CreateUUID(), u.Email, u.UserId, time.Now()).
		Scan(&s.Id, &s.Uuid, &s.Email, &s.UserId, &s.CreatedAt)
	return
}

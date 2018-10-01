package main

import (
	"time"
)

type Todo struct {
	Id        int
	UserId    int
	Content   string
	Completed int
	Due       time.Time
	CreatedAt time.Time
}

func (t *Todo) createTodo() (err error) {
	err = Db.QueryRow("INSERT into todos (user_id, content, completed, due, created_at) values ($1, $2, $3, $4, $5) returning id", t.UserId, t.Content, t.Completed, t.Due, t.CreatedAt).Scan(&t.Id)
	return
}

func (t *Todo) putContent(c string) {
	Db.QueryRow("UPDATE todos SET content = $1 WHERE user_id = $2", c, t.UserId)
	return
}

func (t *Todo) putCompleted(c int) {
	Db.QueryRow("UPDATE todos SET completed = $1 WHERE user_id = $2", c, t.UserId)
	return
}

func (t *Todo) putDue(ti time.Time) {
	Db.QueryRow("UPDATE todos SET due = $1 WHERE user_id = $2", ti, t.UserId)
	return
}

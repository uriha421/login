package main

import (
	"time"
)

type Todo struct {
	Id        int
	UserId    int
	Body      string
	Completed int
	Due       time.Time
	CreatedAt time.Time
}

func (t *Todo) createTodo() (err error) {
	err = Db.QueryRow("INSERT into todos (user_id, body, completed, due, created_at) values ($1, $2, $3, $4, $5) returning id", t.UserId, t.Body, t.Completed, t.Due, t.CreatedAt).Scan(&t.Id)
	return
}

func (t *Todo) putContent(b string) {
	Db.QueryRow("UPDATE todos SET body = $1 WHERE user_id = $2", b, t.UserId)
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

func getTodos(u int) (ts []Todo, err error) {
	rows, err := Db.Query("SELECT (body, completed, due) from todos WHERE user_id = $1", u)
	if err != nil {
		return
	}

	for rows.Next() {
		t := Todo{}
		err = rows.Scan(&t.Body, &t.Completed, &t.Due)
		if err != nil {
			return
		}
		ts = append(ts, t)
	}
	return
}

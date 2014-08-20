package models

import "time"

type User struct {
	Id        int64
	Email     int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func AllUsers() []User {
	users := []User{}
	db.Find(&users)
	return users
}

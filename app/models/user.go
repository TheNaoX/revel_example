package models

import "time"

type User struct {
	Id                int64
	Email             string
	EncryptedPassword []byte
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func AllUsers() []User {
	users := []User{}
	db.Find(&users)
	return users
}

func CreateUser(user *User) {
	db.Save(user)
}

func FindUserById(value string) User {
	user := User{}
	db.Where("id = ?", value).First(&user)
	return user
}

func FindUserByEmail(value string) User {
	user := User{}
	db.Where("email = ?", value).First(&user)
	return user
}

package services

import (
	"revel_example/app/models"

	"code.google.com/p/go.crypto/bcrypt"
)

func AuthenticateUser(email string, password string, success func(int64), err func(string)) {
	user := models.FindUserByEmail(email)
	if user.Id != 0 {
		if PasswordValid(user.EncryptedPassword, password) {
			success(user.Id)
		} else {
			err("Invalid login credentials")
		}
	} else {
		err("Invalid login credentials")
	}
}

func PasswordValid(user_password []byte, password string) bool {
	result := bcrypt.CompareHashAndPassword(user_password, []byte(password))
	return result == nil
}

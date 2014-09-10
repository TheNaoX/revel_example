package services

import (
	"revel_example/app/models"

	"code.google.com/p/go.crypto/bcrypt"
)

func RegisterUser(user models.User, password string, password_confirmation string, success func(int64), err func(string)) {
	if password == password_confirmation {
		if !UserExists(user) {
			encrypted_password, _ := bcrypt.GenerateFromPassword([]byte(password), 5)
			user.EncryptedPassword = encrypted_password
			models.CreateUser(&user)
			success(user.Id)
		} else {
			err("User alredy exists")
		}
	} else {
		err("User alredy exists")
	}
}

func UserExists(user models.User) bool {
	u := models.FindUserByEmail(user.Email)
	return u.Id != 0
}

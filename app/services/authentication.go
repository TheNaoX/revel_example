package services

import (
	"revel_example/app/models"

	"code.google.com/p/go.crypto/bcrypt"

	"github.com/revel/revel"
	"strconv"
)

func AuthenticateUser(context *revel.Controller, user models.User, password string, password_confirmation string) {
	if password == password_confirmation {
		if !UserExists(user) {
			encrypted_password, _ := bcrypt.GenerateFromPassword([]byte(password), 1)
			user.EncryptedPassword = encrypted_password
			models.CreateUser(&user)
			context.Session["user_id"] = strconv.FormatInt(user.Id, 10)
		} else {
			context.Flash.Error("User alredy exists")
		}
	} else {
		context.Flash.Error("Passwords don't match.")
	}
}

func UserExists(user models.User) bool {
	u := models.FindUserByEmail(user.Email)
	return u.Id != 0
}

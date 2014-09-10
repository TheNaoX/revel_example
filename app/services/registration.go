package services

import (
	"encoding/base64"
	"revel_example/app/models"

	"crypto/rand"

	"code.google.com/p/go.crypto/bcrypt"
)

func RegisterUser(user models.User, password string, password_confirmation string, success func(string), err func(string)) {
	if password == password_confirmation {
		if !UserExists(user) {
			CreateUserFromRegistration(user, password, success)
		} else {
			err("User alredy exists")
		}
	} else {
		err("User alredy exists")
	}
}

func CreateUserFromRegistration(user models.User, password string, success func(string)) {
	encrypted_password, _ := bcrypt.GenerateFromPassword([]byte(password), 5)
	user.EncryptedPassword = encrypted_password
	user.AuthToken = GenerateAuthToken()

	models.CreateUser(&user)
	success(user.AuthToken)
}

func GenerateAuthToken() string {
	var random_string string
	for {
		random_string, _ = GenerateRandomString(10)
		u := models.FindUserByAuthToken(random_string)
		if u.Id == 0 {
			break
		}
	}
	return random_string
}

func UserExists(user models.User) bool {
	u := models.FindUserByEmail(user.Email)
	return u.Id != 0
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

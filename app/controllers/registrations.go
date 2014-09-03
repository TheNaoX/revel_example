package controllers

import (
	"revel_example/app/models"
	"revel_example/app/routes"
	"revel_example/app/services"
	"strconv"

	"github.com/revel/revel"
)

type Registrations struct {
	Application
}

func (c Registrations) New() revel.Result {
	if c.UserSignedIn() {
		return c.Redirect(routes.Persons.Index())
	}
	return c.Render()
}

func (c Registrations) Create(user models.User, password string, password_confirmation string) revel.Result {
	services.AuthenticateUser(user, password, password_confirmation, func(id int64) {
		c.Session["user_id"] = strconv.FormatInt(user.Id, 10)
	}, func(message string) {
		c.Flash.Error(message)
	})
	return c.Redirect(routes.Persons.Index())
}

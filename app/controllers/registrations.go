package controllers

import (
	"revel_example/app/models"
	"revel_example/app/routes"
	"revel_example/app/services"

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
	services.RegisterUser(user, password, password_confirmation, func(token string) {
		c.Session["session_token"] = token
	}, func(message string) {
		c.Flash.Error(message)
	})
	return c.Redirect(routes.Persons.Index())
}

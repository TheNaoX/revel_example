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
	return c.Render()
}

func (c Registrations) Create(user models.User, password string, password_confirmation string) revel.Result {
	services.AuthenticateUser(c.Controller, user, password, password_confirmation)
	return c.Redirect(routes.Registrations.New())
}

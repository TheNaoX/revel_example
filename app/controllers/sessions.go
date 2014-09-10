package controllers

import (
	"revel_example/app/routes"
	"revel_example/app/services"
	"strconv"

	"github.com/revel/revel"
)

type Sessions struct {
	Application
}

func (c Sessions) New() revel.Result {
	return c.Render()
}

func (c Sessions) Create(email string, password string) revel.Result {
	var redirect_to string
	services.AuthenticateUser(email, password, func(id int64) {
		c.Session["user_id"] = strconv.FormatInt(id, 10)
		redirect_to = routes.Persons.Index()
	}, func(message string) {
		c.Flash.Error(message)
		redirect_to = routes.Sessions.New()
	})

	return c.Redirect(redirect_to)
}

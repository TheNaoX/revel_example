package controllers

import (
	"revel_example/app/models"
	"revel_example/app/routes"
	"revel_example/app/services"
	"github.com/revel/revel"
	"strconv"
	"fmt"
)

type Registrations struct {
	Application
}

func (c Registrations) New() revel.Result {
	fmt.Println(c.Session["user_id"])
	if c.CurrentUser.Id == 0 {
		return c.Render()
	}

	return c.Redirect(routes.Persons.Index())
}

func (c Registrations) Create(user models.User, password string, password_confirmation string) revel.Result {
	services.AuthenticateUser(user, password, password_confirmation, func(id int64){
		c.Session["user_id"] = strconv.FormatInt(user.Id, 10)
	}, func(message string) {
		c.Flash.Error(message)
	})
	return c.Redirect(routes.Registrations.New())
}

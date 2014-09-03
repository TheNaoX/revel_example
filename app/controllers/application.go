package controllers

import (
	"revel_example/app/models"
	"github.com/revel/revel"
)

type Application struct {
	*revel.Controller
}

func (c Application) GetUserSession() models.User {
	user := models.FindUserById(c.Session["user_id"])
	return user
}

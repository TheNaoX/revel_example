package controllers

import (
	"revel_example/app/models"

	"github.com/revel/revel"
)

type Application struct {
	*revel.Controller
	CurrentUser models.User
}

func (c Application) UserSignedIn() bool {
	c.CurrentUser = models.FindUserById(c.Session["user_id"])
	return c.CurrentUser.Id != 0
}

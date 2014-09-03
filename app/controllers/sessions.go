package controllers

import (
	"revel_example/app/routes"

	"github.com/revel/revel"
)

type Sessions struct {
	Application
}

func (c Sessions) New() revel.Result {
	return c.Render()
}

func (c Sessions) Create() revel.Result {
	return c.Redirect(routes.Persons.Index())
}

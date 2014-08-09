package controllers

import (
	"revel_example/app/models"
	"revel_example/app/routes"

	"github.com/revel/revel"
)

type Persons struct {
	*revel.Controller
}

func (c Persons) Index() revel.Result {
	persons := models.AllPersons()
	return c.Render(persons)
}

func (c Persons) Show(id int64) revel.Result {
	person := models.FindPerson(id)
	return c.Render(person)
}

func (c Persons) New() revel.Result {

	return c.Render()
}

func (c Persons) Create(person models.Person) revel.Result {
	models.CreatePerson(person)
	return c.Redirect(routes.Persons.Index())
}

func (c Persons) Edit(id int64) revel.Result {
	person := models.FindPerson(id)
	return c.Render(person)
}

func (c Persons) Update(id int64, person models.Person) revel.Result {
	models.UpdatePerson(id, person)
	return c.Redirect(routes.Persons.Index())
}

func (c Persons) Delete(id int64) revel.Result {
	models.DeletePerson(id)
	return c.Redirect(routes.Persons.Index())
}

package controllers

import (
	"revel_example/app/models"
	"revel_example/app/routes"

	"github.com/revel/revel"
)

type Persons struct {
	Application
}

func (c Persons) Index() revel.Result {
	persons := models.AllPersons()
	users := models.AllUsers()
	return c.Render(persons, users)
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
	person.Id = id
	person.Update()
	return c.Redirect(routes.Persons.Index())
}

func (c Persons) Delete(id int64) revel.Result {
	person := models.FindPerson(id)
	person.Delete()
	return c.Redirect(routes.Persons.Index())
}

package controllers

import (
	"revel_example/app/models"
	"revel_example/app/routes"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/revel/revel"
)

func SetupDB() gorm.DB {
	db, err := gorm.Open("postgres", "dbname=martini_example sslmode=disable")
	db.LogMode(true)
	PanicIf(err)
	return db
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

type Persons struct {
	*revel.Controller
}

var db gorm.DB = SetupDB()

func (c Persons) Index() revel.Result {
	persons := []models.Person{}
	db.Find(&persons)

	return c.Render(persons)
}

func (c Persons) Show(id int) revel.Result {
	person := models.Person{}
	db.First(&person, c.Params.Get("id"))
	return c.Render(person)
}

func (c Persons) New() revel.Result {

	return c.Render()
}

func (c Persons) Create(person models.Person) revel.Result {
	db.Create(&person)
	return c.Redirect(routes.Persons.Index())
}

func (c Persons) Edit() revel.Result {
	person := models.Person{}
	db.First(&person, c.Params.Get("id"))
	return c.Render(person)
}

func (c Persons) Update(id int64, person models.Person) revel.Result {
	person.Id = id
	db.Save(&person)
	return c.Redirect(routes.Persons.Index())
}

func (c Persons) Delete() revel.Result {
	person := models.Person{}
	db.First(&person, c.Params.Get("id"))
	db.Delete(&person)
	return c.Redirect(routes.Persons.Index())
}

package models

import (
	"revel_example/conf"
	"time"

	"github.com/jinzhu/gorm"
)

type Person struct {
	Id        int64
	Name      string
	Age       int
	Email     string
	Job       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

var db gorm.DB = conf.SetupDB(Person{})

func AllPersons() []Person {
	persons := []Person{}
	db.Find(&persons)

	return persons
}

func FindPerson(id int64) Person {
	person := Person{}
	db.First(&person, id)

	return person
}

func CreatePerson(person Person) Person {
	db.Create(&person)
	return person
}

func UpdatePerson(id int64, person Person) Person {
	person.Id = id
	db.Save(&person)
	return person
}

func DeletePerson(id int64) {
	person := Person{}
	db.First(&person, id)
	db.Delete(&person)
}

package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
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

var db gorm.DB = SetupDB()

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

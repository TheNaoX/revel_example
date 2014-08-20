package models

import (
	"revel_example/conf"

	_ "time"

	"github.com/jinzhu/gorm"
)

var db gorm.DB = conf.SetupDB(func(db gorm.DB) {
	// Migrations!!
	db.AutoMigrate(User{})
	db.AutoMigrate(Person{})
})

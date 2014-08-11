package conf

import (
	"fmt"
	"reflect"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Model interface{}

func SetupDB(model Model) gorm.DB {
	db, err := gorm.Open("postgres", "dbname=martini_example sslmode=disable")
	db.LogMode(true)
	fmt.Println("Auto-Migrating", reflect.TypeOf(model))
	db.AutoMigrate(model)
	PanicIf(err)
	return db
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

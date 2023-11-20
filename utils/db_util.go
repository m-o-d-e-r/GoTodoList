package utils

import (
	"path"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var databaseObject *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open(path.Join(PROGRAM_HOME, "go-todo-list.db")), &gorm.Config{})
	if err != nil {
		panic("failed to initialize database")
	}

	databaseObject = db
}

func GetDB() *gorm.DB {
	return databaseObject
}

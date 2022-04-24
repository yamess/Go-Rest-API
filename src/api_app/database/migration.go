package database

import (
	"fmt"
	"github.com/api_app/constants"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitialMigration() {
	db, err := gorm.Open(postgres.Open(constants.PGdbURL), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}

	dbSQL, ok := db.DB()

	if ok != nil {
		defer dbSQL.Close()
	}
}

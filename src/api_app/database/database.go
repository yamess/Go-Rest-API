package database

import (
	"fmt"
	"github.com/api_app/constants"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreSQL struct {
	Conn *gorm.DB
}

var PG = PostgreSQL{}

func (pg *PostgreSQL) Connect() {
	db, err := gorm.Open(postgres.Open(constants.PGdbURL), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}

	pg.Conn = db
}

func AutoMigrate(models ...interface{}) {

	PG.Connect()

	for _, v := range models {
		ok := PG.Conn.AutoMigrate(&v)
		if ok != nil {
			panic(ok)
		}
	}

}

package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DBCon *gorm.DB
)

func InitDB() {
	var err error

	dsn := "host=localhost user=postgres password=admin dbname=globofactory"
	DBCon, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		panic(err)
		panic("failed to connect to database")
	}

}

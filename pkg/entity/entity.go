package entity

import (
	"fmt"
	"github.com/scotthaley/globofactory/internal/database"
)

type Entity struct {
	Code        string `gorm:"primaryKey"`
	Display     string
}

func SearchEntityTypes(search string) []Entity {
	var entities []Entity
	database.DBCon.Where("code LIKE ?", fmt.Sprintf("%%%v%%", search)).Find(&entities)
	return entities
}

func GetEntity(code string) Entity {
	var e Entity
	database.DBCon.First(&e, "code = ?", code)
	return e
}
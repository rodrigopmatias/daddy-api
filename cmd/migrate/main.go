package main

import (
	"log"

	"github.com/rodrigopmatias/daddy-api/db/models"
	"github.com/rodrigopmatias/daddy-api/helpers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config := helpers.GetConfig()
	db, err := gorm.Open(mysql.Open(config.DbDSN), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(
		&models.Terminal{},
		&models.Metric{},
	)
}

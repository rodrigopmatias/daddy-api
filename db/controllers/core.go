package controllers

import (
	"net/http"

	"github.com/rodrigopmatias/daddy-api/helpers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var config = helpers.GetConfig()

type _CoreController struct{}

func (_CoreController) openConnection() (*gorm.DB, *ControllerError) {
	db, err := gorm.Open(mysql.Open(config.DbDSN), &gorm.Config{})
	if err != nil {
		return nil, NewControllerError(err.Error(), http.StatusInternalServerError)
	}

	return db, nil
}

package controllers

import (
	"net/http"

	"github.com/rodrigopmatias/daddy-api/helpers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Filter struct {
	Expression string
	Value      interface{}
}

func NewFilter(expression string, value interface{}) Filter {
	return Filter{Expression: expression, Value: value}
}

var config = helpers.GetConfig()

type _CoreController struct{}

func (_CoreController) openConnection() (*gorm.DB, *ControllerError) {
	db, err := gorm.Open(mysql.Open(config.DbDSN), &gorm.Config{})
	if err != nil {
		return nil, NewControllerError(err.Error(), http.StatusInternalServerError)
	}

	return db, nil
}

func (_CoreController) ApplyFilters(db *gorm.DB, filters ...Filter) *gorm.DB {
	for _, filter := range filters {
		db = db.Where(filter.Expression, filter.Value)
	}

	return db
}

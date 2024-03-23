package controllers

import (
	"net/http"
	"net/url"
	"time"

	"github.com/rodrigopmatias/daddy-api/helpers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type Filter struct {
	Expression string
	Value      interface{}
}

func NewFilter(expression string, value interface{}) Filter {
	return Filter{Expression: expression, Value: value}
}

type ExtractFilterFunc func(key string, value string) Filter
type ExtractFilterMap map[string]ExtractFilterFunc

func ExtractFilters(baseURL *url.URL, filterMap ExtractFilterMap) []Filter {
	filters := make([]Filter, 0)

	for attr := range baseURL.Query() {
		extractFilter, ok := filterMap[attr]
		if ok {
			filters = append(filters, extractFilter(attr, baseURL.Query().Get(attr)))
		}
	}

	return filters
}

var config = helpers.GetConfig()

type _CoreController struct {
	db *gorm.DB
}

func (c *_CoreController) openConnection() (*gorm.DB, *ControllerError) {
	if c.db == nil {
		db, err := gorm.Open(mysql.New(mysql.Config{DSN: config.DbDSN}))
		if err != nil {
			return nil, NewControllerError(err.Error(), http.StatusInternalServerError)
		}

		db.Use(
			dbresolver.Register(dbresolver.Config{}).
				SetConnMaxIdleTime(time.Minute * 5).
				SetConnMaxLifetime(time.Minute * 15).
				SetMaxIdleConns(5).
				SetMaxOpenConns(15),
		)

		c.db = db
	}

	return c.db, nil
}

func (_CoreController) applyOrder(db *gorm.DB, orders ...string) *gorm.DB {
	for _, order := range orders {
		db = db.Order(order)
	}

	return db
}

func (_CoreController) applyFilters(db *gorm.DB, filters ...Filter) *gorm.DB {
	for _, filter := range filters {
		db = db.Where(filter.Expression, filter.Value)
	}

	return db
}

package controllers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/rodrigopmatias/daddy-api/db/input"
	"github.com/rodrigopmatias/daddy-api/db/models"
)

type _MetricController struct {
	_CoreController
}

func (c _MetricController) Exists(filters ...Filter) (bool, *ControllerError) {
	db, err := c.openConnection()
	if err != nil {
		return false, err
	}

	var count int64
	tx := c.ApplyFilters(db, filters...).Select("id").Limit(1).Model(&models.Metric{}).Count(&count)
	if tx.Error != nil {
		return false, NewControllerError(tx.Error.Error(), http.StatusUnprocessableEntity)
	}

	return count > 0, nil
}

func (c _MetricController) Count(filters ...Filter) (int64, *ControllerError) {
	db, err := c.openConnection()
	if err != nil {
		return 0, err
	}

	var count int64
	tx := c.ApplyFilters(db, filters...).Select("id").Model(&models.Metric{}).Count(&count)
	if tx.Error != nil {
		return 0, NewControllerError(tx.Error.Error(), http.StatusUnprocessableEntity)
	}

	return count, nil
}

func (c _MetricController) Update(id string, values interface{}) *ControllerError {
	db, err := c.openConnection()
	if err != nil {
		return err
	}

	tx := db.Where("id = ?", id).Updates(values).Model(&models.Metric{})
	if tx.Error != nil {
		return NewControllerError(tx.Error.Error(), http.StatusUnprocessableEntity)
	}

	return nil
}

func (c _MetricController) List(offset int, limit int) ([]models.Metric, *ControllerError) {
	db, err := c.openConnection()
	if err != nil {
		return nil, err
	}

	var items []models.Metric
	tx := db.Offset(offset).Limit(limit).Find(&items)
	if tx.Error != nil {
		return nil, NewControllerError(tx.Error.Error(), http.StatusInternalServerError)
	}

	return items, nil
}

func (c _MetricController) Delete(id string) *ControllerError {
	db, err := c.openConnection()
	if err != nil {
		return err
	}

	tx := db.Where("id = ?", id).Delete(&models.Metric{})
	if tx.Error != nil {
		return NewControllerError(tx.Error.Error(), http.StatusUnprocessableEntity)
	}

	if tx.RowsAffected > 1 {
		tx.Rollback()
		return NewControllerError("bad selection to delete", http.StatusBadRequest)
	}

	if tx.RowsAffected == 0 {
		return NewControllerError(fmt.Sprintf("metric with id (%s) not found", id), http.StatusBadRequest)
	}

	return nil
}

func (c _MetricController) Get(id string) (*models.Metric, *ControllerError) {
	db, err := c.openConnection()
	if err != nil {
		return nil, err
	}

	metric := models.Metric{}
	tx := db.Take(&metric, "id = ?", id)
	if tx.Error != nil {
		return nil, NewControllerError(tx.Error.Error(), http.StatusUnprocessableEntity)
	}

	if metric.Id == "" {
		return nil, NewControllerError(fmt.Sprintf("metric with id \"%s\" not found", id), http.StatusNotFound)
	}

	return &metric, nil
}

func (c _MetricController) Create(data input.Metric) (*models.Metric, *ControllerError) {
	return c.CreateWithId(uuid.NewString(), data)
}

func (c _MetricController) CreateWithId(id string, data input.Metric) (*models.Metric, *ControllerError) {
	if err := data.IsValid(); err != nil {
		return nil, NewControllerError(err.Error(), http.StatusUnprocessableEntity)
	}

	db, err := c.openConnection()
	if err != nil {
		return nil, err
	}

	metric := models.Metric{
		Id:         id,
		TerminalId: data.TerminalId,
		CreatedAt:  data.CreatedAt,
	}

	tx := db.Create(&metric)
	if tx.Error != nil {
		return nil, NewControllerError(tx.Error.Error(), http.StatusUnprocessableEntity)
	}

	return &metric, nil
}

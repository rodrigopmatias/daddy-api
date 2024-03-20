package controllers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/rodrigopmatias/daddy-api/db/input"
	"github.com/rodrigopmatias/daddy-api/db/models"
)

type _TerminalController struct {
	_CoreController
}

func (c _TerminalController) Exists(filters ...Filter) (bool, *ControllerError) {
	db, err := c.openConnection()
	if err != nil {
		return false, err
	}

	var count int64
	tx := c.ApplyFilters(db, filters...).Select("id").Limit(1).Model(&models.Terminal{}).Count(&count)
	if tx.Error != nil {
		return false, NewControllerError(tx.Error.Error(), http.StatusUnprocessableEntity)
	}

	return count > 0, nil
}

func (c _TerminalController) Count() (int64, *ControllerError) {
	db, err := c.openConnection()
	if err != nil {
		return 0, err
	}

	var count int64
	tx := db.Select("id").Model(&models.Terminal{}).Count(&count)
	if tx.Error != nil {
		return 0, NewControllerError(tx.Error.Error(), http.StatusUnprocessableEntity)
	}

	return count, nil
}

func (c _TerminalController) Update(id string, values interface{}) *ControllerError {
	db, err := c.openConnection()
	if err != nil {
		return err
	}

	tx := db.Where("id = ?", id).Updates(values).Model(&models.Terminal{})
	if tx.Error != nil {
		return NewControllerError(tx.Error.Error(), http.StatusUnprocessableEntity)
	}

	return nil
}

func (c _TerminalController) List(offset int, limit int) ([]models.Terminal, *ControllerError) {
	db, err := c.openConnection()
	if err != nil {
		return nil, err
	}

	var items []models.Terminal
	tx := db.Offset(offset).Limit(limit).Find(&items)
	if tx.Error != nil {
		return nil, NewControllerError(tx.Error.Error(), http.StatusInternalServerError)
	}

	return items, nil
}

func (c _TerminalController) Delete(id string) *ControllerError {
	db, err := c.openConnection()
	if err != nil {
		return err
	}

	tx := db.Where("id = ?", id).Delete(&models.Terminal{})
	if tx.Error != nil {
		return NewControllerError(tx.Error.Error(), http.StatusUnprocessableEntity)
	}

	if tx.RowsAffected > 1 {
		tx.Rollback()
		return NewControllerError("bad selection to delete", http.StatusBadRequest)
	}

	if tx.RowsAffected == 0 {
		return NewControllerError(fmt.Sprintf("terminal with id (%s) not found", id), http.StatusBadRequest)
	}

	return nil
}

func (c _TerminalController) Get(id string) (*models.Terminal, *ControllerError) {
	db, err := c.openConnection()
	if err != nil {
		return nil, err
	}

	terminal := models.Terminal{}
	tx := db.Take(&terminal, "id = ?", id)
	if tx.Error != nil {
		return nil, NewControllerError(tx.Error.Error(), http.StatusUnprocessableEntity)
	}

	if terminal.Id == "" {
		return nil, NewControllerError(fmt.Sprintf("terminal with id \"%s\" not found", id), http.StatusNotFound)
	}

	return &terminal, nil
}

func (c _TerminalController) Create(data input.Terminal) (*models.Terminal, *ControllerError) {
	return c.CreateWithId(uuid.NewString(), data)
}

func (c _TerminalController) CreateWithId(id string, data input.Terminal) (*models.Terminal, *ControllerError) {
	if err := data.IsValid(); err != nil {
		return nil, NewControllerError(err.Error(), http.StatusUnprocessableEntity)
	}

	db, err := c.openConnection()
	if err != nil {
		return nil, err
	}

	terminal := models.Terminal{
		Id:   id,
		Name: data.Name,
	}

	tx := db.Create(&terminal)
	if tx.Error != nil {
		return nil, NewControllerError(tx.Error.Error(), http.StatusUnprocessableEntity)
	}

	return &terminal, nil
}

package postgres

import (
	"assignment2/model"
	"assignment2/repository"

	"gorm.io/gorm"
)

type orders struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) repository.OrdersRepo {
	return &orders{db}
}

func (ord *orders) Create(param model.Order) (model.Order, error) {
	err := ord.db.Create(&param).Error
	return param, err
}

func (ord *orders) Get() ([]model.Order, error) {
	var ords []model.Order
	err := ord.db.Find(&ords).Error
	return ords, err
}

func (ord *orders) Update(param model.Order) (model.Order, error) {
	err := ord.db.Save(&param).Error
	return param, err
}

func (ord *orders) Delete(id int64) error {
	return ord.db.Delete(&model.Order{}, id).Error
}

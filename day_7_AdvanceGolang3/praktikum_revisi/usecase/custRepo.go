package usecase

import (
	"custommer-crud/infrastructure/config"
	"custommer-crud/model"
)

type CustRepo interface {
	CreateCustomer(model.Custommer) (model.Custommer, error)
	GetById(int) (model.Custommer, error)
	GetAll() (model.Custommers, error)
	UpdateCust(model.Custommer) (model.Custommer, error)
	DeleteCust(model.Custommer) error
}

type CustRepository struct {
	DbHandler config.DbHandler
}

func (db *CustRepository) CreateCustomer(c model.Custommer) (cust model.Custommer, err error) {
	if err = db.DbHandler.Create(&c).Error; err != nil {
		return
	}
	cust = c
	return
}

func (db *CustRepository) GetById(id int) (cust model.Custommer, err error) {
	if err = db.DbHandler.Find(&cust, id).Error; err != nil {
		return
	}
	return
}

func (db *CustRepository) GetAll() (cust model.Custommers, err error) {
	if err = db.DbHandler.Find(&cust).Error; err != nil {
		return
	}
	return
}

func (db *CustRepository) UpdateCust(c model.Custommer) (cust model.Custommer, err error) {
	if err = db.DbHandler.Save(&c).Error; err != nil {
		return
	}
	return
}

func (db *CustRepository) DeleteCust(c model.Custommer) (err error) {
	if err = db.DbHandler.Delete(&c).Error; err != nil {
		return
	}
	return
}

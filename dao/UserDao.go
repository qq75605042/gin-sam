package dao

import (
	"gin-sam/datasource"
	"gin-sam/models"
)

func GetAll() (result []*models.User, err error) {
	err = datasource.InstanceDb().Find(&result).Error
	return
}

func GetByName(name string) (result models.User, err error) {
	err = datasource.InstanceDb().Where(&models.User{
		Account: name,
	}).First(&result).Error
	return
}

package dao

import (
	"gin-sam/models"

	"github.com/jinzhu/gorm"
)

type UserDayNumDao struct {
	db *gorm.DB
}

func (d *UserDayNumDao) GetAll() (result []*models.UserDayNum, err error) {
	err = d.db.Find(&result).Error
	return
}

func (d *UserDayNumDao) Get(id int) (result *models.UserDayNum, err error) {
	err = d.db.Where("id = ", id).First(&result).Error
	return
}

func (d *UserDayNumDao) Delete(id int) error {
	return d.db.Where("id = ", id).Delete(models.UserDayNum{}).Error
}
func (d *UserDayNumDao) CountAll() (count int64, err error) {
	err = d.db.Model(&models.UserDayNum{}).Count(&count).Error
	return
}

func (d *UserDayNumDao) Update(id int, data interface{}) error {
	return d.db.Model(&models.UserDayNum{}).Where("id = ", id).Updates(data).Error
}

func (d *UserDayNumDao) Create(data *models.UserDayNum) error {
	return d.db.Create(&data).Error
}

func NewUserDayNumDao(db *gorm.DB) *UserDayNumDao {
	return &UserDayNumDao{db}
}

package dao

import (
	"gin-sam/models"

	"github.com/jinzhu/gorm"
)

type UserDayNumDao struct {
	db *gorm.DB
}

func (d *UserDayNumDao) GetAll() ([]*models.UserDayNum, error) {
	var result []*models.UserDayNum
	err := d.db.Find(&result).Error
	return result, err
}

func (d *UserDayNumDao) Get(id int) (*models.UserDayNum, error) {
	var result models.UserDayNum
	err := d.db.Where("id = ?", id).First(&result).Error
	return &result, err
}

func (d *UserDayNumDao) Delete(id int) error {
	return d.db.Where("id = ?", id).Delete(models.UserDayNum{}).Error
}
func (d *UserDayNumDao) CountAll() (int64, error) {
	var count int64
	err := d.db.Model(&models.UserDayNum{}).Count(&count).Error
	return count, err
}

func (d *UserDayNumDao) Update(id int, data interface{}) error {
	return d.db.Model(&models.UserDayNum{}).Where("id = ?", id).Updates(data).Error
}

func (d *UserDayNumDao) Create(data *models.UserDayNum) error {
	return d.db.Create(&data).Error
}

func NewUserDayNumDao(db *gorm.DB) *UserDayNumDao {
	return &UserDayNumDao{db}
}

package dao

import (
	"gin-sam/models"

	"github.com/jinzhu/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

func (d *UserDao) GetAll() (result []*models.User, err error) {
	err = d.db.Find(&result).Error
	return
}

func (d *UserDao) Get(id int) (result *models.User, err error) {
	err = d.db.Where("id = ", id).First(&result).Error
	return
}

func (d *UserDao) CountAll() (count int64, err error) {
	err = d.db.Model(&models.User{}).Count(&count).Error
	return
}

func (d *UserDao) Delete(id int) error {
	return d.db.Where("id = ?", id).Delete(models.User{}).Error
}

func (d *UserDao) Update(id int, data interface{}) error {
	return d.db.Model(&models.User{}).Where("id = ", id).Updates(data).Error
}

func (d *UserDao) Create(data *models.User) error {
	return d.db.Create(&data).Error
}

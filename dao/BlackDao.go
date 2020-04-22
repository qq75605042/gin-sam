package dao

import (
	"gin-sam/models"

	"github.com/jinzhu/gorm"
)

type BlackDao struct {
	db *gorm.DB
}

func (d *BlackDao) GetAll() (result []*models.Black, err error) {
	err = d.db.Find(&result).Error
	return
}

func (d *BlackDao) Get(id int) (result *models.Black, err error) {
	err = d.db.Where("id = ", id).First(&result).Error
	return
}

func (d *BlackDao) Delete(id int) error {
	return d.db.Where("id = ", id).Delete(models.Black{}).Error
}
func (d *BlackDao) CountAll() (count int64, err error) {
	err = d.db.Model(&models.Black{}).Count(&count).Error
	return
}

func (d *BlackDao) Update(id int, data interface{}) error {
	return d.db.Model(&models.Black{}).Where("id = ", id).Updates(data).Error
}

func (d *BlackDao) Create(data *models.Black) error {
	return d.db.Create(&data).Error
}

func NewBlackDao(db *gorm.DB) *BlackDao {
	return &BlackDao{db}
}

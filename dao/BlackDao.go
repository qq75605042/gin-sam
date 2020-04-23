package dao

import (
	"gin-sam/models"

	"github.com/jinzhu/gorm"
)

type BlackDao struct {
	db *gorm.DB
}

func (d *BlackDao) GetAll() ([]*models.Black, error) {
	var result []*models.Black
	err := d.db.Find(&result).Error
	return result, err
}

func (d *BlackDao) Get(id int) (*models.Black, error) {
	var result models.Black
	err := d.db.Where("id = ?", id).First(&result).Error
	return &result, err
}

func (d *BlackDao) Delete(id int) error {
	return d.db.Where("id = ?", id).Delete(models.Black{}).Error
}
func (d *BlackDao) CountAll() (int64, error) {
	var count int64
	err := d.db.Model(&models.Black{}).Count(&count).Error
	return count, err
}

func (d *BlackDao) Update(id int, data interface{}) error {
	return d.db.Model(&models.Black{}).Where("id = ?", id).Updates(data).Error
}

func (d *BlackDao) Create(data *models.Black) error {
	return d.db.Create(&data).Error
}

func NewBlackDao(db *gorm.DB) *BlackDao {
	return &BlackDao{db}
}

package dao

import (
	"gin-sam/models"

	"github.com/jinzhu/gorm"
)

type BlackIpDao struct {
	db *gorm.DB
}

func (d *BlackIpDao) GetAll() ([]*models.Blackip, error) {
	var result []*models.Blackip
	err := d.db.Find(&result).Error
	return result, err
}

func (d *BlackIpDao) Get(id int) (*models.Blackip, error) {
	var result models.Blackip
	err := d.db.Where("id = ?", id).First(&result).Error
	return &result, err
}

func (d *BlackIpDao) Delete(id int) error {
	return d.db.Where("id = ?", id).Delete(models.Blackip{}).Error
}
func (d *BlackIpDao) CountAll() (int64, error) {
	var count int64
	err := d.db.Model(&models.Blackip{}).Count(&count).Error
	return count, err
}

func (d *BlackIpDao) Update(id int, data interface{}) error {
	return d.db.Model(&models.Blackip{}).Where("id = ?", id).Updates(data).Error
}

func (d *BlackIpDao) Create(data *models.Blackip) error {
	return d.db.Create(&data).Error
}

func NewBlackIpDao(db *gorm.DB) *BlackIpDao {
	return &BlackIpDao{db}
}

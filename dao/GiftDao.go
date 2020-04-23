package dao

import (
	"gin-sam/models"

	"github.com/jinzhu/gorm"
)

type GiftDao struct {
	db *gorm.DB
}

func (d *GiftDao) GetAll() ([]*models.Gift, error) {
	var result []*models.Gift
	err := d.db.Find(&result).Error
	return result, err
}

func (d *GiftDao) Get(id int) (*models.Gift, error) {
	var result models.Gift
	err := d.db.Where("id = ?", id).First(&result).Error
	return &result, err
}

func (d *GiftDao) Delete(id int) error {
	return d.db.Where("id = ?", id).Delete(models.Gift{}).Error
}
func (d *GiftDao) CountAll() (int64, error) {
	var count int64
	err := d.db.Model(&models.Gift{}).Count(&count).Error
	return count, err
}

func (d *GiftDao) Update(id int, data interface{}) error {
	return d.db.Model(&models.Gift{}).Where("id = ?", id).Updates(data).Error
}

func (d *GiftDao) Create(data *models.Gift) error {
	return d.db.Create(&data).Error
}

func NewGiftDao(db *gorm.DB) *GiftDao {
	return &GiftDao{db}
}

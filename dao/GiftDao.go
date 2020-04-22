package dao

import (
	"gin-sam/models"

	"github.com/jinzhu/gorm"
)

type GiftDao struct {
	db *gorm.DB
}

func (d *GiftDao) GetAll() (result []*models.Gift, err error) {
	err = d.db.Find(&result).Error
	return
}

func (d *GiftDao) Get(id int) (result *models.Gift, err error) {
	err = d.db.Where("id = ", id).First(&result).Error
	return
}

func (d *GiftDao) Delete(id int) error {
	return d.db.Where("id = ", id).Delete(models.Gift{}).Error
}
func (d *GiftDao) CountAll() (count int64, err error) {
	err = d.db.Model(&models.Gift{}).Count(&count).Error
	return
}

func (d *GiftDao) Update(id int, data interface{}) error {
	return d.db.Model(&models.Gift{}).Where("id = ", id).Updates(data).Error
}

func (d *GiftDao) Create(data *models.Gift) error {
	return d.db.Create(&data).Error
}

func NewGiftDao(db *gorm.DB) *GiftDao {
	return &GiftDao{db}
}

package dao

import (
	"gin-sam/models"

	"github.com/jinzhu/gorm"
)

type BlackIpDao struct {
	db *gorm.DB
}

func (d *BlackIpDao) GetAll() (result []*models.Blackip, err error) {
	err = d.db.Find(&result).Error
	return
}

func (d *BlackIpDao) Get(id int) (result *models.Blackip, err error) {
	err = d.db.Where("id = ", id).First(&result).Error
	return
}

func (d *BlackIpDao) Delete(id int) error {
	return d.db.Where("id = ", id).Delete(models.Blackip{}).Error
}
func (d *BlackIpDao) CountAll() (count int64, err error) {
	err = d.db.Model(&models.Blackip{}).Count(&count).Error
	return
}

func (d *BlackIpDao) Update(id int, data interface{}) error {
	return d.db.Model(&models.Blackip{}).Where("id = ", id).Updates(data).Error
}

func (d *BlackIpDao) Create(data *models.Blackip) error {
	return d.db.Create(&data).Error
}

func NewBlackIpDao(db *gorm.DB) *BlackIpDao {
	return &BlackIpDao{db}
}

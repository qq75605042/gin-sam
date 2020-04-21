package dao

import (
	"gin-sam/models"

	"github.com/jinzhu/gorm"
)

type CouponDao struct {
	db *gorm.DB
}

func (d *CouponDao) GetAll() (result []*models.Coupon, err error) {
	err = d.db.Find(&result).Error
	return
}

func (d *CouponDao) Get(id int) (result *models.Coupon, err error) {
	err = d.db.Where("id = ", id).First(&result).Error
	return
}

func (d *CouponDao) Delete(id int) error {
	return d.db.Where("id = ", id).Delete(models.Coupon{}).Error
}
func (d *CouponDao) CountAll() (count int64, err error) {
	err = d.db.Model(&models.Coupon{}).Count(&count).Error
	return
}

func (d *CouponDao) Update(id int, data interface{}) error {
	return d.db.Model(&models.Coupon{}).Where("id = ", id).Updates(data).Error
}

func (d *CouponDao) Create(data *models.Coupon) error {
	return d.db.Create(&data).Error
}

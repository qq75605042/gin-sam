package dao

import (
	"gin-sam/models"

	"github.com/jinzhu/gorm"
)

type CouponDao struct {
	db *gorm.DB
}

func (d *CouponDao) GetAll() ([]*models.Coupon, error) {
	var result []*models.Coupon
	err := d.db.Find(&result).Error
	return result, err
}

func (d *CouponDao) Get(id int) (*models.Coupon, error) {
	var result models.Coupon
	err := d.db.Where("id = ?", id).First(&result).Error
	return &result, err
}

func (d *CouponDao) Delete(id int) error {
	return d.db.Where("id = ?", id).Delete(models.Coupon{}).Error
}
func (d *CouponDao) CountAll() (int64, error) {
	var count int64
	err := d.db.Model(&models.Coupon{}).Count(&count).Error
	return count, err
}

func (d *CouponDao) Update(id int, data interface{}) error {
	return d.db.Model(&models.Coupon{}).Where("id = ?", id).Updates(data).Error
}

func (d *CouponDao) Create(data *models.Coupon) error {
	return d.db.Create(&data).Error
}

func NewCouponDao(db *gorm.DB) *CouponDao {
	return &CouponDao{db}
}

package services

import (
	"gin-sam/dao"
	"gin-sam/datasource"
	"gin-sam/models"
)

type CouponService interface {
	GetAll() ([]*models.Coupon, error)
	CountAll() (int64, error)
	Get(id int) (*models.Coupon, error)
	Delete(id int) error
	Update(id int, data interface{}) error
	Create(data *models.Coupon) error
}

type couponService struct {
	dao *dao.CouponDao
}

func (u *couponService) GetAll() ([]*models.Coupon, error) {
	return u.dao.GetAll()
}

func (u *couponService) CountAll() (int64, error) {
	return u.dao.CountAll()
}

func (u *couponService) Get(id int) (*models.Coupon, error) {
	return u.dao.Get(id)
}

func (u *couponService) Delete(id int) error {
	return u.dao.Delete(id)
}

func (u *couponService) Update(id int, data interface{}) error {
	return u.dao.Update(id, data)
}

func (u *couponService) Create(data *models.Coupon) error {
	return u.dao.Create(data)
}

func NewCouponService() CouponService {
	return &couponService{dao: dao.NewCouponDao(datasource.InstanceDb())}
}

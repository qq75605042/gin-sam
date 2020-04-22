package services

import (
	"gin-sam/dao"
	"gin-sam/datasource"
	"gin-sam/models"
)

type GiftService interface {
	GetAll() ([]*models.Gift, error)
	CountAll() (int64, error)
	Get(id int) (*models.Gift, error)
	Delete(id int) error
	Update(id int, data interface{}) error
	Create(data *models.Gift) error
}

type giftService struct {
	dao *dao.GiftDao
}

func (u *giftService) GetAll() ([]*models.Gift, error) {
	return u.dao.GetAll()
}

func (u *giftService) CountAll() (int64, error) {
	return u.dao.CountAll()
}

func (u *giftService) Get(id int) (*models.Gift, error) {
	return u.dao.Get(id)
}

func (u *giftService) Delete(id int) error {
	return u.dao.Delete(id)
}

func (u *giftService) Update(id int, data interface{}) error {
	return u.dao.Update(id, data)
}

func (u *giftService) Create(data *models.Gift) error {
	return u.dao.Create(data)
}

func NewGiftService() GiftService {
	return &giftService{dao: dao.NewGiftDao(datasource.InstanceDb())}
}

package services

import (
	"gin-sam/dao"
	"gin-sam/datasource"
	"gin-sam/models"
)

type UserDayNumService interface {
	GetAll() ([]*models.Gift, error)
	CountAll() (int64, error)
	Get(id int) (*models.Gift, error)
	Delete(id int) error
	Update(id int, data interface{}) error
	Create(data *models.Gift) error
}

type userDayNumService struct {
	dao *dao.GiftDao
}

func (u *userDayNumService) GetAll() ([]*models.Gift, error) {
	return u.dao.GetAll()
}

func (u *userDayNumService) CountAll() (int64, error) {
	return u.dao.CountAll()
}

func (u *userDayNumService) Get(id int) (*models.Gift, error) {
	return u.dao.Get(id)
}

func (u *userDayNumService) Delete(id int) error {
	return u.dao.Delete(id)
}

func (u *userDayNumService) Update(id int, data interface{}) error {
	return u.dao.Update(id, data)
}

func (u *userDayNumService) Create(data *models.Gift) error {
	return u.dao.Create(data)
}

func NewUserDayNumService() UserDayNumService {
	return &userDayNumService{dao: dao.NewGiftDao(datasource.InstanceDb())}
}

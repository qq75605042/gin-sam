package services

import (
	"gin-sam/dao"
	"gin-sam/datasource"
	"gin-sam/models"
)

type BlackIpService interface {
	GetAll() ([]*models.Blackip, error)
	CountAll() (int64, error)
	Get(id int) (*models.Blackip, error)
	Delete(id int) error
	Update(id int, data interface{}) error
	Create(data *models.Blackip) error
}

type blackIpService struct {
	dao *dao.BlackIpDao
}

func (u *blackIpService) GetAll() ([]*models.Blackip, error) {
	return u.dao.GetAll()
}

func (u *blackIpService) CountAll() (int64, error) {
	return u.dao.CountAll()
}

func (u *blackIpService) Get(id int) (*models.Blackip, error) {
	return u.dao.Get(id)
}

func (u *blackIpService) Delete(id int) error {
	return u.dao.Delete(id)
}

func (u *blackIpService) Update(id int, data interface{}) error {
	return u.dao.Update(id, data)
}

func (u *blackIpService) Create(data *models.Blackip) error {
	return u.dao.Create(data)
}

func NewBlackIpService() BlackIpService {
	return &blackIpService{dao: dao.NewBlackIpDao(datasource.InstanceDb())}
}

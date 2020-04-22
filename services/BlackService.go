package services

import (
	"gin-sam/dao"
	"gin-sam/datasource"
	"gin-sam/models"
)

type BlackService interface {
	GetAll() ([]*models.Black, error)
	CountAll() (int64, error)
	Get(id int) (*models.Black, error)
	Delete(id int) error
	Update(id int, data interface{}) error
	Create(data *models.Black) error
}

type blackService struct {
	dao *dao.BlackDao
}

func (u *blackService) GetAll() ([]*models.Black, error) {
	return u.dao.GetAll()
}

func (u *blackService) CountAll() (int64, error) {
	return u.dao.CountAll()
}

func (u *blackService) Get(id int) (*models.Black, error) {
	return u.dao.Get(id)
}

func (u *blackService) Delete(id int) error {
	return u.dao.Delete(id)
}

func (u *blackService) Update(id int, data interface{}) error {
	return u.dao.Update(id, data)
}

func (u *blackService) Create(data *models.Black) error {
	return u.dao.Create(data)
}

func NewBlackService() BlackService {
	return &blackService{dao: dao.NewBlackDao(datasource.InstanceDb())}
}

package services

import (
	"gin-sam/dao"
	"gin-sam/datasource"
	"gin-sam/models"
)

type UserService interface {
	GetAll() ([]*models.User, error)
	CountAll() (int64, error)
	Get(id int) (*models.User, error)
	Delete(id int) error
	Update(id int, data interface{}) error
	Create(data *models.User) error
}

type userService struct {
	dao *dao.UserDao
}

func (u *userService) GetAll() ([]*models.User, error) {
	return u.dao.GetAll()
}

func (u *userService) CountAll() (int64, error) {
	return u.dao.CountAll()
}

func (u *userService) Get(id int) (*models.User, error) {
	return u.dao.Get(id)
}

func (u *userService) Delete(id int) error {
	return u.dao.Delete(id)
}

func (u *userService) Update(id int, data interface{}) error {
	return u.dao.Update(id, data)
}

func (u *userService) Create(data *models.User) error {
	return u.dao.Create(data)
}

func NewUserService() UserService {
	return &userService{dao: dao.NewUserDao(datasource.InstanceDb())}
}

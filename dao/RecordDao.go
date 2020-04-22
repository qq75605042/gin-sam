package dao

import (
	"gin-sam/models"

	"github.com/jinzhu/gorm"
)

type RecordDao struct {
	db *gorm.DB
}

func (d *RecordDao) GetAll() (result []*models.Records, err error) {
	err = d.db.Find(&result).Error
	return
}

func (d *RecordDao) Get(id int) (result *models.Records, err error) {
	err = d.db.Where("id = ", id).First(&result).Error
	return
}

func (d *RecordDao) Delete(id int) error {
	return d.db.Where("id = ", id).Delete(models.Records{}).Error
}
func (d *RecordDao) CountAll() (count int64, err error) {
	err = d.db.Model(&models.Records{}).Count(&count).Error
	return
}

func (d *RecordDao) Update(id int, data interface{}) error {
	return d.db.Model(&models.Records{}).Where("id = ", id).Updates(data).Error
}

func (d *RecordDao) Create(data *models.Records) error {
	return d.db.Create(&data).Error
}

func NewRecordDao(db *gorm.DB) *RecordDao {
	return &RecordDao{db}
}

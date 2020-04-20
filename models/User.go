package models

import "github.com/jinzhu/gorm"

type Users struct {
	gorm.Model
	Account   string `gorm:"column:account;type:varchar(100);not null" json:"account"`     // 账户
	Password  string `gorm:"column:password;type:varchar(255);not null" json:"password"`   // 密码
	SysStatus int8   `gorm:"column:sys_status;type:tinyint(1);not null" json:"sys_status"` // 0正常 1删除
}

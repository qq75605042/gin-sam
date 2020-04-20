package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Account string	`json:"account"`
	Password string `json:"password"`
	SysStatus uint `json:"sysStatus"`
}

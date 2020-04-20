package models

type Blackip struct {
	ID         int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	IP         string `gorm:"column:ip;type:varchar(255);not null" json:"ip"`
	Blacktime  int    `gorm:"column:blacktime;type:int(10);not null" json:"blacktime"` // 黑名单到期时间
	SysCreated int    `gorm:"column:sys_created;type:int(10)" json:"sys_created"`
	SysUpdated int    `gorm:"column:sys_updated;type:int(10)" json:"sys_updated"`
}
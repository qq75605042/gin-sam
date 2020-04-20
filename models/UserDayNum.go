package models

type UserDayNum struct {
	ID         int `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	UId        int `gorm:"column:uid;type:int(10);not null" json:"uid"` // 用户
	Day        int `gorm:"column:day;type:int(10);not null" json:"day"` // 某一天，202001
	Num        int `gorm:"column:num;type:int(1);not null" json:"num"`  // 次数
	SysCreated int `gorm:"column:sys_created;type:int(10)" json:"sys_created"`
	SysUpdated int `gorm:"column:sys_updated;type:int(10)" json:"sys_updated"`
}

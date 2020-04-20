package models

type Black struct {
	ID         int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	Username   string `gorm:"column:username;type:varchar(50);not null" json:"username"` // 用户名
	Blacktime  int    `gorm:"column:blacktime;type:int(10);not null" json:"blacktime"`   // 黑名单到期时间
	Realname   string `gorm:"column:realname;type:varchar(50);not null" json:"realname"`
	Mobile     string `gorm:"column:mobile;type:varchar(50);not null" json:"mobile"`
	Address    string `gorm:"column:address;type:varchar(255)" json:"address"`
	SysIP      string `gorm:"column:sys_ip;type:varchar(255)" json:"sys_ip"`
	SysCreated int    `gorm:"column:sys_created;type:int(10)" json:"sys_created"`
	SysUpdated int    `gorm:"column:sys_updated;type:int(10)" json:"sys_updated"`
}
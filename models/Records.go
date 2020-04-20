package models

type Records struct {
	ID         int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	GiftID     int    `gorm:"column:gift_id;type:int(10);not null" json:"gift_id"`          // 奖品id
	GiftName   string `gorm:"column:gift_name;type:varchar(255);not null" json:"gift_name"` // 奖品名称
	GiftType   int8   `gorm:"column:gift_type;type:tinyint(1);not null" json:"gift_type"`   // 奖品类型
	UId        int    `gorm:"column:uid;type:int(10);not null" json:"uid"`                  // 用户
	UserName   string `gorm:"column:user_name;type:varchar(255);not null" json:"user_name"` // 用户名
	PrizeCode  int    `gorm:"column:prize_code;type:int(1);not null" json:"prize_code"`     // 中奖编码，4位随机数
	GiftData   string `gorm:"column:gift_data;type:varchar(255);not null" json:"gift_data"` // 奖品信息
	SysCreated int    `gorm:"column:sys_created;type:int(10);not null" json:"sys_created"`
	SysStatus  int8   `gorm:"column:sys_status;type:tinyint(1)" json:"sys_status"` // 0正常 1删除 2作弊
	SysIP      string `gorm:"column:sys_ip;type:varchar(255)" json:"sys_ip"`
}

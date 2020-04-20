package models
type Coupon struct {
	ID         int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	GiftID     int    `gorm:"column:gift_id;type:int(11);not null" json:"gift_id"` // 奖品id
	Code       string `gorm:"column:code;type:varchar(255);not null" json:"code"`  // 编号
	SysStatus  int8   `gorm:"column:sys_status;type:tinyint(1)" json:"sys_status"`
	SysCreated int    `gorm:"column:sys_created;type:int(10);not null" json:"sys_created"`
	SysUpdated int    `gorm:"column:sys_updated;type:int(10)" json:"sys_updated"`
}
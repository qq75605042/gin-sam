package models
type Gift struct {
	ID           int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	Title        string `gorm:"column:title;type:varchar(255);not null" json:"title"`               // 奖品名称
	PrizeNum     int    `gorm:"column:prize_num;type:int(10) unsigned;not null" json:"prize_num"`   // 数量,0无限 >0 有限制 <0无奖品
	LeftNum      int    `gorm:"column:left_num;type:int(10) unsigned;not null" json:"left_num"`     // 剩余数量
	PrizeCode    string `gorm:"column:prize_code;type:varchar(50);not null" json:"prize_code"`      // 0-9999，中奖区间0-0表示万分之一
	PrizeTime    int    `gorm:"column:prize_time;type:int(11) unsigned;not null" json:"prize_time"` // 发奖周期，天
	Img          string `gorm:"column:img;type:varchar(255);not null" json:"img"`                   // 奖品图片
	Displayorder int    `gorm:"column:displayorder;type:int(255);not null" json:"displayorder"`     // 越小越靠前
	Gtype        int8   `gorm:"column:gtype;type:tinyint(1) unsigned;not null" json:"gtype"`        // 奖品类型 0虚拟币 1券 2实物小奖 3实物大奖
	Gdata        string `gorm:"column:gdata;type:varchar(255)" json:"gdata"`                        // 扩展数据，如虚拟币数量
	TimeBegin    int    `gorm:"column:time_begin;type:int(10);not null" json:"time_begin"`          // 开始时间
	TimeEnd      int    `gorm:"column:time_end;type:int(10);not null" json:"time_end"`              // 结束时间
	PrizeData    string `gorm:"column:prize_data;type:mediumtext" json:"prize_data"`                // 发奖计划，[[time1,num1]...]
	PrizeBegin   int    `gorm:"column:prize_begin;type:int(10);not null" json:"prize_begin"`        // 发奖周期开始
	PrizeEnd     int    `gorm:"column:prize_end;type:int(10);not null" json:"prize_end"`            // 发奖周期结束
	SysStatus    int8   `gorm:"column:sys_status;type:tinyint(1);not null" json:"sys_status"`       // 0正常 1删除
	SysCreated   int    `gorm:"column:sys_created;type:int(10);not null" json:"sys_created"`        // 创建时间
	SysUpdated   int    `gorm:"column:sys_updated;type:int(10);not null" json:"sys_updated"`        // 更新时间
	SysIP        string `gorm:"column:sys_ip;type:varchar(50);not null" json:"sys_ip"`              // 操作人
}
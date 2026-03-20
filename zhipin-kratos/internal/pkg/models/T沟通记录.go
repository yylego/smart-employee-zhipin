package models

import (
	"time"

	"gorm.io/gorm"
)

// E事件类型 is the database enum type for communication event type
type E事件类型 string

const (
	C事件类型_类型未知 E事件类型 = "类型未知"
	C事件类型_发消息   E事件类型 = "发消息"
	C事件类型_收消息   E事件类型 = "收消息"
	C事件类型_发简历   E事件类型 = "发简历"
	C事件类型_开聊限制 E事件类型 = "开聊限制"
	C事件类型_安排面试 E事件类型 = "安排面试"
	C事件类型_收到邀请 E事件类型 = "收到邀请"
	C事件类型_被拒绝   E事件类型 = "被拒绝"
)

// T沟通记录 records each interaction event with a recruiter — one record per event.
//
// T沟通记录 记录与招聘者的每次交互事件——每个事件一条记录。
type T沟通记录 struct {
	gorm.Model
	P岗位主键 uint      `gorm:"column:position_id;index:idx_communications_position_id;not null"` // 关联岗位ID
	E事件类型 E事件类型 `gorm:"column:event_type;type:varchar(32);not null"` // 事件类型
	E事件时间 time.Time `gorm:"column:event_time;not null"`                  // 事件发生时间
	C消息内容 string    `gorm:"column:content;type:text"`                    // 消息内容或备注
	D消息方向 int32     `gorm:"column:direction"`                            // 0=我方发出, 1=对方发出
}

func (T沟通记录) TableName() string {
	return "communications"
}

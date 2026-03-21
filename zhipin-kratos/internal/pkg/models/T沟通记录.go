package models

import "gorm.io/gorm"

// T沟通记录 stores each chat message — full conversation synced per position.
//
// T沟通记录 存储每条聊天消息——按岗位全量同步整个对话。
type T沟通记录 struct {
	gorm.Model
	P岗位主键  uint   `gorm:"column:position_id;index:idx_communications_position_id;not null"` // 关联岗位ID
	D消息方向  int32  `gorm:"column:direction;not null"`                                        // 0=我方 1=对方
	C消息内容  string `gorm:"column:content;type:text;not null"`                                // 消息文本
	T消息时间  int64  `gorm:"column:timestamp;not null"`                                        // unix 秒
	B简历消息  bool   `gorm:"column:is_resume"`                                                 // 是否为简历消息
	R简历版本  string `gorm:"column:resume_version;type:varchar(256)"`                           // 简历文件名
}

func (T沟通记录) TableName() string {
	return "communications"
}

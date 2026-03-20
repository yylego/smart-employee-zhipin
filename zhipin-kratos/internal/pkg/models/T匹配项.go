package models

import "gorm.io/gorm"

// E匹配状态 is the database enum type for match status
type E匹配状态 string

const (
	C匹配状态_状态未知 E匹配状态 = "状态未知"
	C匹配状态_匹配     E匹配状态 = "匹配"
	C匹配状态_部分匹配 E匹配状态 = "部分匹配"
	C匹配状态_不匹配   E匹配状态 = "不匹配"
)

// T匹配项 is the requirement-level match analysis — one record per requirement per position.
//
// T匹配项 是需求级别的匹配分析——每个岗位的每条要求对应一条记录。
type T匹配项 struct {
	gorm.Model
	P岗位主键 uint     `gorm:"column:position_id;index:idx_match_items_position_id;not null"` // 关联岗位ID
	R岗位要求 string   `gorm:"column:requirement;type:text;not null"`        // 岗位的一条要求
	M匹配状态 E匹配状态 `gorm:"column:match_status;type:varchar(32);not null"` // 匹配状态
	R简历对应 string   `gorm:"column:resume_point;type:text"`                // 简历中对应的经验
	R补充说明 string   `gorm:"column:remark;type:varchar(512)"`              // 补充说明
	S排序序号 int32    `gorm:"column:sort_index"`                            // 显示顺序
}

func (T匹配项) TableName() string {
	return "match_items"
}

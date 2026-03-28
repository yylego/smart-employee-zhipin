package models

import "gorm.io/gorm"

// E岗位状态 is the database enum type for position status
type E岗位状态 string

const (
	C岗位状态_状态未知 E岗位状态 = "状态未知"
	C岗位状态_待处理  E岗位状态 = "待处理"
	C岗位状态_已跳过  E岗位状态 = "已跳过"
	C岗位状态_开聊限制 E岗位状态 = "开聊限制"
	C岗位状态_已发消息 E岗位状态 = "已发消息"
	C岗位状态_已回复  E岗位状态 = "已回复"
	C岗位状态_已发简历 E岗位状态 = "已发简历"
	C岗位状态_面试中  E岗位状态 = "面试中"
	C岗位状态_已拿到  E岗位状态 = "已拿到"
	C岗位状态_已拒绝  E岗位状态 = "已拒绝"
	C岗位状态_不再联系 E岗位状态 = "不再联系"
)

// C岗位活跃状态 groups statuses that represent active/in-progress positions
var C岗位活跃状态 = []E岗位状态{C岗位状态_已发消息, C岗位状态_已回复, C岗位状态_已发简历}

// T岗位 is the core entity — one record per job posting
// T岗位 是核心实体——每条招聘岗位一条记录
type T岗位 struct {
	gorm.Model
	J岗位编号 string `gorm:"column:job_id;type:varchar(28);uniqueIndex:idx_positions_job_id;not null"` // BOSS直聘28位岗位ID
	T岗位名称 string `gorm:"column:title;type:varchar(256);not null"`                                  // 岗位名称
	C公司名称 string `gorm:"column:company;type:varchar(256);not null"`                                // 公司名称
	S薪资范围 string `gorm:"column:salary_range;type:varchar(64);not null"`                            // e.g. "35-65K·15薪"
	S薪资下限 int32  `gorm:"column:salary_min"`                                                        // 月薪下限（K）
	S薪资上限 int32  `gorm:"column:salary_max"`                                                        // 月薪上限（K）
	C城市名称 string `gorm:"column:city;type:varchar(64);not null"`                                    // 城市
	L岗位链接 string `gorm:"column:link;type:varchar(512)"`                                            // 完整岗位链接
	R招聘者  string `gorm:"column:recruiter;type:varchar(128)"`                                       // 招聘者姓名和职位
	E招聘者号 string `gorm:"column:enc_boss_id;type:varchar(64)"`                                      // encBossId
	I猎头标记 bool   `gorm:"column:is_hunter"`                                                         // 是否猎头岗位
	S岗位状态 E岗位状态  `gorm:"column:status;type:varchar(32)"`                                           // 岗位状态
	S跳过原因 string `gorm:"column:skip_reason;type:text"`                                             // 跳过原因
	M匹配度  int32  `gorm:"column:match_rate"`                                                        // 0-100匹配度百分比
	N备注信息 string `gorm:"column:notes;type:text"`                                                   // 备注
	L最后沟通 int64  `gorm:"column:last_comm_at"`                                                      // 最后沟通时间（unix秒）
	L最后方向 int32  `gorm:"column:last_comm_dir"`                                                     // 最后消息方向 0=我方 1=对方
	L简历版本 string `gorm:"column:last_resume;type:varchar(256)"`                                     // 最后发的简历版本文件名
	D岗位职责 string `gorm:"column:duties;type:text"`                                                  // 岗位职责
	R岗位要求 string `gorm:"column:requirements;type:text"`                                            // 岗位要求
}

func (T岗位) TableName() string {
	return "positions"
}
